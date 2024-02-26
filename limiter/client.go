package limiter

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//go:generate counterfeiter -generate

//counterfeiter:generate -o mocks/rate_limiter.go --fake-name RateLimiter . RateLimiterInterface
type RateLimiterInterface interface {
	RoundTrip(req *http.Request) (*http.Response, error)
}

type roundTripper struct {
	Transport http.RoundTripper
}

func NewHttpClient(rt http.RoundTripper) *http.Client {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &http.Client{
		Transport: &roundTripper{
			Transport: rt,
		},
	}
}

func (c *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	response, err := c.Transport.RoundTrip(req)
	if err != nil || !isRateLimitExceeded(response) {
		return response, err
	}

	sleepDuration := getSleepDuration(response)
	fmt.Printf("Rate limit exceeded, sleeping for %0.0f minutes\n", sleepDuration.Minutes())
	time.Sleep(sleepDuration)
	return c.Transport.RoundTrip(req)
}

func isRateLimitExceeded(response *http.Response) bool {
	rateLimitRemaining, _ := strconv.ParseInt(response.Header.Get("X-Ratelimit-Remaining"), 10, 64)
	return rateLimitRemaining == 0
}

func getSleepDuration(response *http.Response) time.Duration {
	rateLimitReset, _ := strconv.ParseInt(response.Header.Get("X-Ratelimit-Reset"), 10, 64)
	resetTime := time.Unix(rateLimitReset, 0)
	return time.Until(resetTime).Round(time.Minute) + 5*time.Second
}
