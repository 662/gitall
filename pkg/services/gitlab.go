package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Res2Map(res *http.Response, err error) (map[string]interface{}, error) {
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resData map[string]interface{}
	if err := json.Unmarshal(body, &resData); err != nil {
		return nil, err
	}
	return resData, nil
}

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

func FindGitlabProject(baseUrl string, name string, token string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/v4/projects/%s?private_token=%s", baseUrl, url.QueryEscape(name), url.QueryEscape(token))
	return Res2Map(http.Get(url))
}

func CreateGitlabMergeRequest(baseUrl string, id string, token string, sourceBranch string, targetBranch string) (map[string]interface{}, error) {
	title := fmt.Sprintf("Merge %s to %s", sourceBranch, targetBranch)
	url := fmt.Sprintf("%s/api/v4/projects/%s/merge_requests?private_token=%s", baseUrl, url.QueryEscape(id), url.QueryEscape(token))
	var reqData = fmt.Sprintf(`{
		"source_branch": "%s",
		"target_branch": "%s",
		"title": "%s"
	}`, sourceBranch, targetBranch, title)
	log.Printf("url: %s", url)
	log.Printf("title: %s", title)
	log.Printf("reqData: %s", reqData)
	return Res2Map(http.Post(url, "application/json", bytes.NewBuffer([]byte(reqData))))
}
