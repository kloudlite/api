package types

import (
	"github.com/kloudlite/api/pkg/egob"
	"github.com/kloudlite/api/pkg/repos"
)

type EnvResources string

const (
	EnvResourcesApps             EnvResources = "apps"
	EnvResourcesConfigs          EnvResources = "configs"
	EnvResourcesSecrets          EnvResources = "secrets"
	EnvResourcesManagedResources EnvResources = "managedResources"
)

type NotificationEnvParams struct {
	EnvName string `json:"envName" graphql:"noinput"`
}

type NotificationClusterParams struct {
	ClusterName string `json:"clusterName" graphql:"noinput"`
}

type NotificationResourceType string
type NotificationType string

const (
	NotificationTypeAlert        NotificationType = "alert"
	NotificationTypeNotification NotificationType = "notification"
)

const (
	NotificationResourceTypeEnvironment NotificationResourceType = "environment"
	NotificationResourceTypeCluster     NotificationResourceType = "cluster"
	NotificationResourceTypeAccount     NotificationResourceType = "account"
)

type Notification struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	ResourceType     NotificationResourceType `json:"resourceType" graphql:"noinput"`
	NotificationType NotificationType         `json:"notificationType" graphql:"noinput"`

	NotificationEnvParams     *NotificationEnvParams     `json:"notificationEnvParams" graphql:"noinput"`
	NotificationClusterParams *NotificationClusterParams `json:"notificationClusterParams" graphql:"noinput"`

	AccountName string `json:"accountName" graphql:"noinput"`
	Read        bool   `json:"read" graphql:"noinput"`
}

func (obj *Notification) ToBytes() ([]byte, error) {
	return egob.Marshal(obj)
}

func (obj *Notification) ParseBytes(data []byte) error {
	return egob.Unmarshal(data, obj)
}

var NotificationIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
