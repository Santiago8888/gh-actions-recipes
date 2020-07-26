package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)

	opt := &github.PullRequestListOptions{State: "closed"}
	prs, _, _ := client.PullRequests.List(ctx, "cypress-io", "cypress", opt)

	fmt.Println(prs)
}
