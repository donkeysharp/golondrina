package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ISSUES_URL_FORMAT  = "https://%s/api/v3/issues?filter=assigned"
	AUTH_HEADER_FORMAT = "token %s"
)

func getIssues(githubHost, token string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf(ISSUES_URL_FORMAT, githubHost)
	authHeader := fmt.Sprintf(AUTH_HEADER_FORMAT, token)

	client := http.Client{}

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Authorization", authHeader)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, errors.New(
			fmt.Sprintf("Error in response status code %d\nWith content%s",
				response.StatusCode,
				string(content)))
	}

	var issues []map[string]interface{}
	err = json.Unmarshal(content, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
