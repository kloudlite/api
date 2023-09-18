package entities

import (
	"kloudlite.io/pkg/repos"
)

const (
	RepoAccessReadOnly  RepoAccess = "read"
	RepoAccessReadWrite RepoAccess = "read_write"
)

type Credential struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	AccountName      string     `json:"accountName" graphql:"noinput"`
	Token            string     `json:"token" graphql:"noinput"`
	Access           RepoAccess `json:"access" graphql:"noinput"`
	Name             string     `json:"name" graphql:"noinput"`
}

var CredentialIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
