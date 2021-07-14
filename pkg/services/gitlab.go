package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func FetchGitlabRepositoriesByGroupName(baseUrl string, groupName string, token string) ([]string, error) {
	url := fmt.Sprintf("%s/api/v4/groups/%s/projects?private_token=%s&per_page=999", baseUrl, url.QueryEscape(groupName), url.QueryEscape(token))
	log.Printf("fetch repositorries url from: %s", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var repositoriesArray []map[string]interface{}
	if err := json.Unmarshal(body, &repositoriesArray); err != nil {
		return nil, err
	}
	var repositoriesLength = len(repositoriesArray)
	log.Printf("repositorries length: %d", repositoriesLength)
	urls := make([]string, repositoriesLength)
	for i, repo := range repositoriesArray {
		url, ok := repo["ssh_url_to_repo"].(string)
		if !ok {
			return nil, errors.New("ssh_url_to_repo is not a string")
		}
		urls[i] = url
	}
	return urls, nil
}
