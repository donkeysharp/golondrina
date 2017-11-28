package github

import (
	"github.com/donkeysharp/golondrina/models"
	"strconv"
)

type Parser struct {
	Host  string
	Token string
}

func (p *Parser) GetNotifications() ([]models.NotificationEvent, error) {
	issues, err := getIssues(p.Host, p.Token)
	if err != nil {
		return nil, err
	}
	var result []models.NotificationEvent

	for _, issue := range issues {
		isPullRequest := issue["pull_request"]
		if isPullRequest != nil {
			result = append(result, models.NotificationEvent{
				InternalId: strconv.Itoa(int(issue["id"].(float64))),
				Title:      issue["title"].(string),
				Url:        issue["html_url"].(string),
				Id:         strconv.Itoa(123),
			})
		}
	}

	return result, nil
}
