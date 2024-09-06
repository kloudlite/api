package entities

import (
	"github.com/kloudlite/api/pkg/repos"
)

type Uses struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
}

var UsesIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "targetNamespace", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
