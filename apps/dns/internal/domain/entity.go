package domain

import "kloudlite.io/pkg/repos"

type AccountCName struct {
	repos.BaseEntity `bson:",inline"`
	AccountId        repos.ID `bson:"accountId" json:"accountId"`
	CName            string   `bson:"cName" json:"cName"`
}

type RegionCName struct {
	repos.BaseEntity `bson:",inline"`
	RegionId         repos.ID `bson:"regionId" json:"regionId"`
	CName            string   `bson:"cName" json:"cName"`
	IsShared         bool     `bson:"isShared" json:"isShared"`
}

type NodeIps struct {
	repos.BaseEntity `bson:",inline"`
	RegionPart       string   `bson:"regionPart" json:"regionPart"`
	ClusterPart      string   `bson:"clusterPart" json:"clusterPart"`
	AccountPart      string   `bson:"accountPart" json:"accountPart"`
	Ips              []string `bson:"ips"`
}

type Site struct {
	repos.BaseEntity `bson:",inline"`
	AccountId        repos.ID `bson:"accountId,omitempty" json:"accountId,omitempty"`
	RegionId         repos.ID `bson:"regionId,omitempty" json:"regionId,omitempty"`
	Domain           string   `bson:"host,omitempty" json:"host,omitempty"`
	Verified         bool     `bson:"verified,omitempty" json:"verified,omitempty"`
}

type Record struct {
	repos.BaseEntity `bson:",inline"`
	Host             string   `bson:"host,omitempty" json:"host,omitempty"`
	Answers          []string `bson:"answers,omitempty" json:"answers,omitempty"`
}

var RecordIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "host", Value: repos.IndexAsc},
			{Key: "type", Value: repos.IndexAsc},
			{Key: "answer", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
var SiteIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "domain", Value: repos.IndexAsc},
		},
	},
}
var NodeIpIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}

var RegionCNameIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "regionId", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "cName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}

var AccountCNameIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "accountId", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "cName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
