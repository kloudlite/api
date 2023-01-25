package entities

import (
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	"kloudlite.io/pkg/repos"
)

type ManagedResourceStatus string

const (
	ManagedResourceStateSyncing = ManagedResourceStatus("sync-in-progress")
	ManagedResourceStateLive    = ManagedResourceStatus("live")
	ManagedResourceStateError   = ManagedResourceStatus("error")
	ManagedResourceStateDown    = ManagedResourceStatus("down")
)

type ManagedResource struct {
	repos.BaseEntity       `bson:",inline"`
	crdsv1.ManagedResource `json:",inline" bson:",inline"`

	//ClusterId        repos.ID              `json:"cluster_id" bson:"cluster_id"`
	//ProjectId        repos.ID              `json:"project_id" bson:"project_id"`
	//Name             string                `json:"name" bson:"name"`
	//Namespace        string                `json:"namespace" bson:"namespace"`
	//ResourceType     ManagedResourceType   `json:"resource_type" bson:"resource_type"`
	//ServiceId        repos.ID              `bson:"service_id" json:"service_id"`
	//Values           map[string]string     `json:"values" bson:"values"`
	//Status           ManagedResourceStatus `json:"status" bson:"status"`
	//Conditions       []metav1.Condition    `json:"conditions" bson:"conditions"`
}

var ManagedResourceIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
			{Key: "metadata.namespace", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
