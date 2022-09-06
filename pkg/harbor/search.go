package harbor

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"kloudlite.io/pkg/errors"
	"kloudlite.io/pkg/repos"
)

// Repository created by pasting json from harbor instance network tab
type Repository struct {
	ArtifactCount int       `json:"artifact_count"`
	CreationTime  time.Time `json:"creation_time"`
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	ProjectId     int       `json:"project_id"`
	PullCount     int       `json:"pull_count"`
	UpdateTime    time.Time `json:"update_time"`
}

func (h *Client) SearchRepositories(ctx context.Context, accountId repos.ID, searchQ string, searchOpts ListOptions) ([]Repository, error) {

	req, err := h.NewAuthzRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/repositories", accountId), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("q", fmt.Sprintf("name=~%s", searchQ))
	q.Add(
		"sort", func() string {
			if searchOpts.Sort == "" {
				return "-id"
			}
			return searchOpts.Sort
		}(),
	)
	q.Add("page", strconv.FormatInt(searchOpts.Page, 10))
	q.Add("page_size", strconv.FormatInt(searchOpts.PageSize, 10))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var results []Repository
		if err := json.Unmarshal(body, &results); err != nil {
			return nil, err
		}
		return results, nil
	}

	return nil, errors.Newf("bad status code (%d) received, with error message, %s", resp.StatusCode, body)
}

type Artifact struct {
	Size int `json:"size"`
	Tags []struct {
		RepositoryId int       `json:"repository_id"`
		Name         string    `json:"name"`
		PushTime     time.Time `json:"push_time"`
		PullTime     time.Time `json:"pull_time"`
		Signed       bool      `json:"signed"`
		Id           int       `json:"id"`
		Immutable    bool      `json:"immutable"`
		ArtifactId   int       `json:"artifact_id"`
	} `json:"tags"`
}

type ImageTag struct {
	Name      string `json:"name"`
	Signed    bool   `json:"signed"`
	Immutable bool   `json:"immutable"`
}

type ListOptions struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
	Sort     string `json:"sort"`
}

type ListTagsOpts struct {
	WithImmutable bool
	WithSignature bool
	ListOptions
}

func (h *Client) ListTags(ctx context.Context, projectName string, repoName string, tagOpts ListTagsOpts) ([]ImageTag, error) {
	req, err := h.NewAuthzRequest(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/repositories/%s/artifacts", projectName, url.PathEscape(repoName)), nil)
	q := req.URL.Query()

	q.Add("with_tag", "true")
	if tagOpts.WithImmutable {
		q.Add("with_immutable", "true")
	}
	if tagOpts.WithSignature {
		q.Add("with_signature", "true")
	}
	q.Add(
		"sort", func() string {
			if tagOpts.Sort == "" {
				return "-id"
			}
			return tagOpts.Sort
		}(),
	)
	q.Add("page", strconv.FormatInt(tagOpts.Page, 10))
	q.Add("page_size", strconv.FormatInt(tagOpts.PageSize, 10))

	req.URL.RawQuery = q.Encode()

	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var artifacts []Artifact
		if err := json.Unmarshal(b, &artifacts); err != nil {
			return nil, err
		}
		tags := make([]ImageTag, 0, len(artifacts))
		for i := range artifacts {
			for j := range artifacts[i].Tags {
				tag := artifacts[i].Tags[j]
				if tag.Name != "" {
					tags = append(tags, ImageTag{Name: tag.Name, Signed: tag.Signed, Immutable: tag.Immutable})
				}
			}
		}
		return tags, nil
	}
	return nil, errors.Newf("bad status code received (%d), with message, %s", resp.StatusCode, b)
}