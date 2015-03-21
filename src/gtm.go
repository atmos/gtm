package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	var client GtmClient

	app := cli.NewApp()
	app.Name = "gtm"
	app.Version = "0.0.1"
	app.Usage = "Manage GitHub teams from the command line"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Usage: "The the user to manage",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "The token for API interaction",
			EnvVar: "GITHUB_TOKEN",
		},
		cli.StringFlag{
			Name:   "server_url",
			Value:  "https://api.github.com",
			Usage:  "The url for the GitHub API",
			EnvVar: "GITHUB_API_URL",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "list",
			Usage:       "list user teams",
			Description: "List the teams the user is a member of.",
			Action: func(c *cli.Context) {
				fmt.Printf("%-v", client)
				println("API URL:", c.GlobalString("server_url"))
				client.Info()
			},
		},
	}
	app.Before = func(c *cli.Context) error {
		client = defaultClient(c.GlobalString("url"),
			c.GlobalString("token"))
		return nil
	}

	app.Run(os.Args)
}
