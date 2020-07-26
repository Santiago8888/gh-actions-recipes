package main

import (
	"fmt"
	"context"
	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	orgs, _, _ := client.Organizations.List(ctx, "willnorris", nil)

	fmt.Println(orgs[0].GetAvatarURL())
}

