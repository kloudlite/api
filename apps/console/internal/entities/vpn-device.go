package entities

import (
	"fmt"

	wireguardV1 "github.com/kloudlite/operator/apis/wireguard/v1"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

type VPNDevice struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`

	wireguardV1.Device `json:",inline" graphql:"uri=k8s://devices.wireguard.kloudlite.io"`

	DisplayName string   `json:"displayName"`
	CreatedBy   repos.ID `json:"createdBy" graphql:"noinput"`
	AccountName string   `json:"accountName" graphql:"noinput"`
	ClusterName string   `json:"clusterName" graphql:"noinput"`

	SyncStatus t.SyncStatus `json:"syncStatus" graphql:"noinput"`
}

var VPNDeviceIndexes = []repos.IndexField{
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

func ValidateVPNDevice(d *VPNDevice) error {
	errMsgs := []string{}
	if d.CreatedBy == "" {
		errMsgs = append(errMsgs, "createdBy is required")
	}

	if d.DisplayName == "" {
		errMsgs = append(errMsgs, "displayName is required")
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf("%v", errMsgs)
	}
	return nil
}