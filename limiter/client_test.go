package limiter_test

import (
	"github.com/lindluni/go-github-rate-limiter/limiter"
	"github.com/lindluni/go-github-rate-limiter/limiter/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestRoundTrip(t *testing.T) {
	rateLimiter := &mocks.RateLimiter{}
	client := limiter.NewHttpClient(rateLimiter)
	rateLimiter.RoundTripReturnsOnCall(0, &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header: http.Header{
			"X-Ratelimit-Remaining": []string{"0"},
			"X-Ratelimit-Reset":     []string{strconv.FormatInt(time.Now().Unix(), 10)},
		},
	}, nil)
	rateLimiter.RoundTripReturnsOnCall(1, &http.Response{
		StatusCode: http.StatusOK,
		Header: http.Header{
			"X-Ratelimit-Remaining": []string{"10"},
		},
	}, nil)

	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRoundTripDefaultTransport(t *testing.T) {
	client := limiter.NewHttpClient(nil)
	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
