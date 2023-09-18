package docker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DockerCli interface {
	ListRepositoryTags(repoName string, limit *int, after *string) ([]string, error)
	DeleteRepositoryTag(repoName string, tag string) error
}

type DockerCliImpl struct {
	registryUrl string
}

func getQuery(limit *int, after *string) (string, error) {

	if after != nil && limit == nil {
		return "", fmt.Errorf("limit must be set if after is set")
	}

	if limit != nil && *limit < 0 {
		return "", fmt.Errorf("limit must be positive")
	}

	if after != nil && *after == "" {
		return "", fmt.Errorf("after must not be empty")
	}

	if limit != nil && after != nil {
		return fmt.Sprintf("?n=%d&last=%s", *limit, *after), nil
	}

	if limit != nil {
		return fmt.Sprintf("?n=%d", *limit), nil
	}

	return "", nil
}

func (d *DockerCliImpl) ListRepositoryTags(repoName string, limit *int, after *string) ([]string, error) {
	s, err := getQuery(limit, after)
	if err != nil {
		return nil, err
	}

	type TagList struct {
		Name string
		Tags []string
	}

	r, err := http.Get(d.registryUrl + "/v2/" + repoName + "/tags/list" + s)
	if err != nil {
		return nil, err
	}

	var tagList TagList
	if err := json.NewDecoder(r.Body).Decode(&tagList); err != nil {
		return nil, err
	}

	return tagList.Tags, nil
}

func (d *DockerCliImpl) DeleteRepositoryTag(repoName string, tag string) error {

	// create a new HTTP client
	client := &http.Client{}

	// create a new DELETE request
	req, err := http.NewRequest(http.MethodDelete, d.registryUrl+"/v2/"+repoName+"/manifests/"+tag, nil)
	if err != nil {
		return err
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func NewDockerCli(registryUrl string) DockerCli {
	return &DockerCliImpl{
		registryUrl: registryUrl,
	}
}
