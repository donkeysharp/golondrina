package cmd

import (
	"github.com/urfave/cli"
)

const (
	APP_NAME    = "Golondrina"
	DESCRIPTION = "Show notifications when a new pull request is assigned to you."
	VERSION     = "0.1"
)

func StartApp(args []string, callback func(string, string, int) error) {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = DESCRIPTION
	app.Version = VERSION

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host,H",
			Value:  "api.github.com",
			Usage:  "Github host (when using Github Enterprise)",
			EnvVar: "GITHUB_HOST",
		},
		cli.StringFlag{
			Name:   "token,t",
			Usage:  "Github token",
			EnvVar: "AUTH_TOKEN",
		},
		cli.IntFlag{
			Name:   "interval,i",
			Value:  10,
			Usage:  "Polling interval in minutes",
			EnvVar: "POLL_INTERVAL",
		},
	}

	app.Action = func(c *cli.Context) error {
		host := c.String("host")
		token := c.String("token")
		interval := c.Int("interval")

		if token == "" {
			return cli.NewExitError("Token must be specified", 1)
		}

		callback(host, token, interval)

		return nil
	}

	app.Run(args)
}
