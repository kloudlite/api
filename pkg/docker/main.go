package docker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Docker interface {
	DeleteDigest(repoName string, tagName string) error
}

type docker struct {
	registryUrl string
}

func (d *docker) DeleteDigest(repoName string, refrence string) error {
	uri := fmt.Sprintf("%s/v2/%s/manifests/%s", d.registryUrl, repoName, refrence)
	req, err := http.NewRequest("DELETE", uri, nil)

	// create a new HTTP client
	client := &http.Client{}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 202 {
		fmt.Println(uri)
		return fmt.Errorf("failed to delete tag %s:%s: %s", repoName, refrence, resp.Status)
	}

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// print the response body
	log.Println(string(body))

	return nil
}

func NewDockerClient(registryUrl string) Docker {
	return &docker{
		registryUrl: registryUrl,
	}
}