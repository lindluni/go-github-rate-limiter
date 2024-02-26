package main

import (
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v59/github"
	"github.com/lindluni/go-github-rate-limiter/limiter"
	"net/http"
	"os"
)

func main() {
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, 841475, 47740245, "key.pem")
	if err != nil {
		panic(err)
	}
	httpClient := limiter.NewHttpClient(itr)
	client := github.NewClient(httpClient).WithAuthToken(os.Getenv("GITHUB_PAT"))
	ctx := context.Background()
	for {
		_, resp, err := client.Repositories.Get(ctx, "google", "go-github")
		if err != nil {
			fmt.Printf("Status: %d\n", resp.StatusCode)
			fmt.Println("5")
			fmt.Println(err.Error())
			break
		}
		response, _, err := client.RateLimit.Get(ctx)
		if err != nil {

			fmt.Println("6")
			panic(err)
		}
		fmt.Println(response.GetCore().Remaining)
	}
}
