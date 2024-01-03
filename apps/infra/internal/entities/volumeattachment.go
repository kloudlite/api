package entities

import (
	"github.com/kloudlite/api/pkg/repos"
	storagev1 "k8s.io/api/storage/v1"
)

type VolumeAttachment struct {
	repos.BaseEntity           `json:",inline" graphql:"noinput"`
	storagev1.VolumeAttachment `json:",inline"`

	AccountName string `json:"accountName" graphql:"noinput"`
	ClusterName string `json:"clusterName" graphql:"noinput"`
}

var VolumeAttachmentIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
			{Key: "accountName", Value: repos.IndexAsc},
			{Key: "clusterName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
