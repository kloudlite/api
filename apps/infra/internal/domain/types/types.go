package types

import (
	"context"
	"time"

	"github.com/kloudlite/api/pkg/repos"
)

type InfraContext struct {
	context.Context
	UserId      repos.ID
	UserEmail   string
	UserName    string
	AccountName string
}

func (i InfraContext) GetUserId() repos.ID {
	return i.UserId
}

func (i InfraContext) GetUserEmail() string {
	return i.UserEmail
}

func (i InfraContext) GetUserName() string {
	return i.UserName
}

type UpdateAndDeleteOpts struct {
	MessageTimestamp time.Time
}

