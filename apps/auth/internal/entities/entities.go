package entities

import (
	"time"

	"github.com/kloudlite/api/pkg/repos"
	"golang.org/x/oauth2"
)

type InvitationStatus string

const (
	InvitationStatusAccepted InvitationStatus = "accepted"
	InvitationStatusRejected InvitationStatus = "rejected"
	InvitationStatusNone     InvitationStatus = "none"
	InvitationStatusSend     InvitationStatus = "sent"
)

type UserMetadata map[string]any

type ProviderDetail struct {
	TokenId repos.ID `json:"token_id" bson:"token_id"`
	Avatar  *string  `json:"avatar" bson:"avatar"`
}

type Session struct {
	ID           repos.ID `json:"id"`
	UserID       repos.ID `json:"user_id"`
	UserEmail    string   `json:"user_email"`
	LoginMethod  string   `json:"login_method"`
	UserVerified bool     `json:"user_verified"`
}

type User struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	Name             string           `json:"name"`
	Avatar           *string          `json:"avatar"`
	ProviderGithub   *ProviderDetail  `json:"provider_github"`
	ProviderGitlab   *ProviderDetail  `json:"provider_gitlab"`
	ProviderGoogle   *ProviderDetail  `json:"provider_google"`
	Email            string           `json:"email"`
	Password         string           `json:"password" graphql:"ignore"`
	InvitationStatus InvitationStatus `json:"invite"`
	Verified         bool             `json:"verified" graphql:"noinput"`
	Metadata         UserMetadata     `json:"metadata"`
	Joined           time.Time        `json:"joined"`
	PasswordSalt     string           `json:"password_salt" graphql:"ignore"`
	Approved         bool             `json:"approved" graphql:"noinput"`
}

var UserIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "email", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}

type AccessToken struct {
	repos.BaseEntity `bson:",inline"`
	UserId           repos.ID       `json:"user_id"`
	Email            string         `json:"email"`
	Provider         string         `json:"provider"`
	Token            *oauth2.Token  `json:"token"`
	Data             map[string]any `json:"data"`
}

var AccessTokenIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "user_id", Value: repos.IndexAsc},
			{Key: "email", Value: repos.IndexAsc},
			{Key: "provider", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}

type VerifyToken struct {
	Token  string   `json:"token"`
	UserId repos.ID `json:"user_id"`
}

type ResetPasswordToken struct {
	Token  string   `json:"token"`
	UserId repos.ID `json:"user_id"`
}

type ChangeEmailToken struct {
	Token  string   `json:"token"`
	UserId repos.ID `json:"user_id"`
}

type LoginStatus string

const (
	LoginStatusPending   = "pending"
	LoginStatusFailed    = "failed"
	LoginStatusSucceeded = "succeeded"
)

type RemoteLogin struct {
	repos.BaseEntity `bson:",inline"`
	LoginStatus      LoginStatus `json:"login_status"`
	Secret           string      `json:"secret"`
	AuthHeader       string      `json:"auth_header"`
}

var RemoteTokenIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}

type InviteCode struct {
	repos.BaseEntity `bson:",inline"`
	Name             string `json:"name"`
	InviteCode       string `json:"inviteCode"`
}

var InviteCodeIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
