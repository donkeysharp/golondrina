package github

import (
	"github.com/donkeysharp/golondrina/store"
	"strconv"
)

type Parser struct {
	Host  string
	Token string
}

func (p *Parser) GetNotifications() ([]store.NotificationEvent, error) {
	issues, err := getIssues(p.Host, p.Token)
	if err != nil {
		return nil, err
	}
	var result []store.NotificationEvent

	for _, issue := range issues {
		isPullRequest := issue["pull_request"]
		if isPullRequest != nil {
			result = append(result, store.NotificationEvent{
				Id:       strconv.Itoa(int(issue["id"].(float64))),
				Title:    issue["title"].(string),
				Url:      issue["html_url"].(string),
				Provider: store.GITHUB_PROVIDER,
			})
		}
	}

	return result, nil
}
