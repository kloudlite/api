package entities

import (
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/repos"
)

type EmailConfig struct {
	Enabled     bool   `json:"enabled"`
	MailAddress string `json:"mailAddress"`
}

type SlackConfig struct {
	Enabled bool   `json:"enabled"`
	Webhook string `json:"webhook"`
}

type TelegramConfig struct {
	Enabled bool `json:"enabled"`
}

type NotificationConf struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	CreatedBy        common.CreatedOrUpdatedBy `json:"createdBy" graphql:"noinput"`
	LastUpdatedBy    common.CreatedOrUpdatedBy `json:"lastUpdatedBy" graphql:"noinput"`

	EmailConfiguration    *EmailConfig    `json:"emailConfigurations"`
	SlackConfiguration    *SlackConfig    `json:"slackConfigurations"`
	TelegramConfiguration *TelegramConfig `json:"telegramConfigurations"`

	AccountName string `json:"accountName" graphql:"noinput"`
}

var NotificationConfIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "accountName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
