package main

import (
	"fmt"
	"github.com/octokit/go-octokit/octokit"
)

type GtmClient struct {
	Client *octokit.Client
}

func defaultClient(url string, token string) GtmClient {
	var client GtmClient

	client.Client = octokit.NewClientWith(url,
		"GitHub Team Membership",
		octokit.TokenAuth{AccessToken: token},
		nil)

	return client
}

func (c *GtmClient) Info() {
	user_url := &octokit.CurrentUserURL

	url, _ := user_url.Expand(nil)

	users, result := c.Client.Users(url).All()

	if result.HasError() {
		fmt.Println(result)
		return
	}

	for _, user := range users {
		fmt.Printf("%v - %s\n", user.ID, user.Login)
	}
}
