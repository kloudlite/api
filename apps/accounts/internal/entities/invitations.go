package entities

import (
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/pkg/repos"
)

type Invitation struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	InvitedBy        string    `json:"invitedBy" graphql:"noinput"`
	InviteToken      string    `json:"inviteToken" graphql:"noinput"`
	UserEmail        string    `json:"userEmail,omitempty"`
	UserName         string    `json:"userName,omitempty"`
	UserRole         iamT.Role `json:"userRole"`
	AccountName      string    `json:"accountName" graphql:"noinput"`
	Accepted         *bool     `json:"accepted,omitempty" graphql:"noinput"`
	Rejected         *bool     `json:"rejected,omitempty" graphql:"noinput"`
}

var InvitationIndices = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "inviteToken", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "accountName", Value: repos.IndexAsc},
			{Key: "userEmail", Value: repos.IndexAsc},
			{Key: "userRole", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
