package main

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/donkeysharp/golondrina/cmd"
	"github.com/donkeysharp/golondrina/github"
	"github.com/donkeysharp/golondrina/store"
	"github.com/robfig/cron"
	"os"
)

func displayNotification(notification store.NotificationEvent) {
	fmt.Println("Displaying notification")
	notify := notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName:     cmd.APP_NAME,
	})

	notify.Push("New Pull Request", notification.Title, "", notificator.UR_NORMAL)
}

func processNotifications(host, token string) {
	fmt.Println("Getting notifications...")

	p := github.Parser{
		Host:  host,
		Token: token,
	}

	notifications, err := p.GetNotifications()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Retrieved %d notifications\n", len(notifications))

	for _, notification := range notifications {
		if !store.NotificationExist(notification) {
			fmt.Println("Notification was not processed before, saving.")
			store.AddNotification(notification)
			displayNotification(notification)
		} else {
			fmt.Println("Notification already processed")
		}
	}
}

func run(host, token string, interval int) error {
	scheduler := cron.New()
	cronExpression := fmt.Sprintf("0 */%d * * * *", interval)
	fmt.Println("Scheduling with cron expression:", cronExpression)

	scheduler.AddFunc(cronExpression, func() { processNotifications(host, token) })
	scheduler.Run()

	return nil
}

func main() {
	cmd.StartApp(os.Args, run)
}
