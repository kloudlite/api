package common

import (
	"encoding/json"
	"github.com/kloudlite/api/pkg/repos"
)

type CreatedOrUpdatedBy struct {
	UserId    repos.ID `json:"userId"`
	UserName  string   `json:"userName"`
	UserEmail string   `json:"userEmail"`
}

const CreatedOnTenantClusterUserName string = "created-on-tenant-cluster-user-name"
const CreatedOnTenantClusterUserId string = "created-on-tenant-cluster-user-id"
const CreatedOnTenantClusterUserEmail string = "created-on-tenant-cluster-user-email"

type ResourceMetadata struct {
	DisplayName string `json:"displayName"`

	CreatedBy     CreatedOrUpdatedBy `json:"createdBy" graphql:"noinput"`
	LastUpdatedBy CreatedOrUpdatedBy `json:"lastUpdatedBy" graphql:"noinput"`
}

type ValidationError struct {
	Label  string
	Errors []string
}

func (v ValidationError) Error() string {
	b, _ := json.Marshal(map[string]any{
		"label":  v.Label,
		"errors": v.Errors,
	})
	return string(b)
}
