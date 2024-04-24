package entities

import (
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	"github.com/kloudlite/operator/pkg/operator"
)

type BluePrintType string

const (
	SingletonBlueprint BluePrintType = "singleton_blueprint"
	GroupBlueprint     BluePrintType = "group_blueprint"
)

type IOTDeviceBlueprint struct {
	repos.BaseEntity        `json:",inline" graphql:"noinput"`
	common.ResourceMetadata `json:",inline"`

	crdsv1.Blueprint `json:",inline"`

	DeploymentName string        `json:"deploymentName"`
	AccountName    string        `json:"accountName" graphql:"noinput"`
	ProjectName    string        `json:"projectName" graphql:"noinput"`
	BluePrintType  BluePrintType `json:"bluePrintType"`

	SyncStatus t.SyncStatus `json:"syncStatus" graphql:"noinput"`
}

func (a *IOTDeviceBlueprint) GetDisplayName() string {
	return a.ResourceMetadata.DisplayName
}

func (a *IOTDeviceBlueprint) GetGeneration() int64 {
	return a.ObjectMeta.Generation
}

func (a *IOTDeviceBlueprint) GetStatus() operator.Status {
	return operator.Status{}
}

var IOTDeviceBlueprintIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: fields.Id, Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{
				Key:   "name",
				Value: repos.IndexAsc,
			},
		},
		Unique: true,
	},
}
