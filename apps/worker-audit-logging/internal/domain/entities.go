package domain

import (
	"github.com/kloudlite/api/pkg/beacon"
	"github.com/kloudlite/api/pkg/repos"
)

type EventLog struct {
	repos.BaseEntity     `json:",inline" bson:",inline"`
	beacon.AuditLogEvent `json:",inline" bson:",inline"`
}

var EventLogIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "who.email", Value: repos.IndexAsc},
		},
	},
	{
		Field: []repos.IndexKey{
			{Key: "who.account_id", Value: repos.IndexAsc},
		},
	},
}
