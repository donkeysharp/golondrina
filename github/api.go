package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	issuesUrlFormat  = "https://%s/api/v3/issues?filter=assigned"
	authHeaderFormat = "token %s"
)

func getIssues(githubHost, token string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf(issuesUrlFormat, githubHost)
	authHeader := fmt.Sprintf(authHeaderFormat, token)

	client := http.Client{}

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Authorization", authHeader)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error 1")
		return nil, err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2")
		return nil, err
	}

	var issues []map[string]interface{}
	err = json.Unmarshal(content, &issues)
	if err != nil {
		fmt.Println("error 3", err)
		return nil, err
	}

	return issues, nil
}
