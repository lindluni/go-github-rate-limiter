# go-github-rate-limiter

`go-github-rate-limiter` is a simple rate limiter built on for the [go-github](https://github.com/google/go-github) 
library. It is designed to allow your application to gracefully hand hitting the 
[GitHub API rate limit](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api).

## Installation

```shell
go get github.com/lindluni/go-github-rate-limiter
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v59/github"
	"github.com/lindluni/go-github-rate-limiter/limiter"
)

func main() {
	httpClient := limiter.NewHttpClient(nil)
	client := github.NewClient(httpClient).WithAuthToken(os.Getenv("GITHUB_PAT"))
	ctx := context.Background()
	repo, resp, err := client.Repositories.Get(ctx, "lindluni", "go-github-rate-limiter")
	if err != nil {
		if resp.StatusCode == 404 {
			fmt.Println("Invalid repository")
		}
		panic(err)
	}
	fmt.Println(repo.GetFullName())
}
```

## Usage with GitHub Apps

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v59/github"
	"github.com/lindluni/go-github-rate-limiter/limiter"
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
	repo, resp, err := client.Repositories.Get(ctx, "lindluni", "go-github-rate-limiter")
	if err != nil {
		if resp.StatusCode == 404 {
			fmt.Println("Invalid repository")
		}
		panic(err)
	}
	fmt.Println(repo.GetFullName())
}
```
