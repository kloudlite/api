package common

import (
	"kloudlite.io/pkg/repos"
)

type CreatedOrUpdatedBy struct {
	UserId    repos.ID `json:"userId"`
	UserName  string   `json:"userName"`
	UserEmail string   `json:"userEmail"`
}

type ResourceMetadata struct {
	DisplayName string `json:"displayName"`

	CreatedBy     CreatedOrUpdatedBy `json:"createdBy" graphql:"noinput"`
	LastUpdatedBy CreatedOrUpdatedBy `json:"updatedBy" graphql:"noinput"`
}
