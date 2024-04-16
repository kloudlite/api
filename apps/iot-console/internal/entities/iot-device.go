package entities

import (
	"fmt"

	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/pkg/repos"
)

type IOTDevice struct {
	repos.BaseEntity        `json:",inline" graphql:"noinput"`
	common.ResourceMetadata `json:",inline"`

	Name           string `json:"name"`
	AccountName    string `json:"accountName" graphql:"noinput"`
	ProjectName    string `json:"projectName" graphql:"noinput"`
	PublicKey      string `json:"publicKey"`
	ServiceCIDR    string `json:"serviceCIDR" graphql:"noinput"`
	PodCIDR        string `json:"podCIDR" graphql:"noinput"`
	IP             string `json:"ip" graphql:"noinput"`
	DeploymentName string `json:"deploymentName" graphql:"noinput"`
	Version        string `json:"version"`

	ClusterToken string `json:"clusterToken" graphql:"noinput"`
	Index        int    `json:"index" graphql:"noinput"`
}

type DeviceWithServices struct {
	*IOTDevice
	ExposedDomains []string `json:"exposedDomains"`
	ExposedIps     []string `json:"exposedIps"`
}

func (i *IOTDevice) GetClusterName() string {
	return fmt.Sprintf("iot-device-%s", i.Name)
}

var IOTDeviceIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: fields.Id, Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "name", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "deploymentName", Value: repos.IndexAsc},
			{Key: "index", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
