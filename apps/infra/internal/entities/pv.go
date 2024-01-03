package entities

import (
	"github.com/kloudlite/api/pkg/repos"
	corev1 "k8s.io/api/core/v1"
)

type PersistentVolume struct {
	repos.BaseEntity        `json:",inline" graphql:"noinput"`
	corev1.PersistentVolume `json:",inline"`

	AccountName string `json:"accountName" graphql:"noinput"`
	ClusterName string `json:"clusterName" graphql:"noinput"`
}

var PersistentVolumeIndices = []repos.IndexField{
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
