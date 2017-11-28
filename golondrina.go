package main

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/donkeysharp/golondrina/cmd"
	"github.com/donkeysharp/golondrina/github"
	"os"
	"time"
)

func foo() {
	notify := notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Golondrina",
	})

	token := os.Getenv("AUTH_TOKEN")
	githubHost := os.Getenv("GITHUB_HOST")
	if githubHost == "" {
		githubHost = "github.com"
	}

	time.Sleep(1 * time.Second)
	notify.Push("New Pull Request", "A new pull request for ReconMVS/recon-app", "", notificator.UR_NORMAL)
	p := github.Parser{
		Host:  "ghe.coxautoinc.com",
		Token: token,
	}
	notifications, _ := p.GetNotifications()
	fmt.Println(notifications)
}

func run(host, token string, interval int) error {
	return nil
}

func main() {
	cmd.StartApp(os.Args, run)
}
