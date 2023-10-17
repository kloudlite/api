package entities

import (
	"kloudlite.io/common"
	"kloudlite.io/pkg/repos"
)

type GitProvider string
type BuildStatus string

const (
	Github GitProvider = "github"
	Gitlab GitProvider = "gitlab"
)

const (
	BuildStatusPending BuildStatus = "pending"
	BuildStatusQueued  BuildStatus = "queued"
	BuildStatusRunning BuildStatus = "running"
	BuildStatusSuccess BuildStatus = "success"
	BuildStatusFailed  BuildStatus = "failed"
	BuildStatusError   BuildStatus = "error"
	BuildStatusIdle    BuildStatus = "idle"
)

type GitSource struct {
	Repository string      `json:"repository"`
	Branch     string      `json:"branch"`
	Provider   GitProvider `json:"provider"`
	WebhookId  *int        `json:"webhookId" graphql:"noinput"`
}

type Build struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	Name        string `json:"name"`
	AccountName string `json:"accountName" graphql:"noinput"`

	CreatedBy     common.CreatedOrUpdatedBy `json:"createdBy" graphql:"noinput"`
	LastUpdatedBy common.CreatedOrUpdatedBy `json:"lastUpdatedBy" graphql:"noinput"`

	Repository string `json:"repository"`
	Tag        string `json:"tag"`

	Source GitSource `json:"source"`

	CredUser common.CreatedOrUpdatedBy `json:"credUser" graphql:"noinput"`

	ErrorMessages map[string]string `json:"errorMessages" graphql:"noinput"`
	Status        BuildStatus       `json:"status" graphql:"noinput"`
}

var BuildIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "repository", Value: repos.IndexAsc},
			{Key: "tag", Value: repos.IndexAsc},
			{Key: "accountName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}