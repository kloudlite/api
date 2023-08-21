package domain

import (
	"context"
	"errors"
	"fmt"
	iamT "kloudlite.io/apps/iam/types"
	"kloudlite.io/grpc-interfaces/kloudlite.io/rpc/iam"
	"kloudlite.io/pkg/repos"

	"go.mongodb.org/mongo-driver/mongo"
)

// access
const (
	READ_PROJECT   = "read_project"
	UPDATE_PROJECT = "update_project"

	READ_ACCOUNT   = "read_account"
	UPDATE_ACCOUNT = "update_account"
)

func mongoError(err error, descp string) error {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(descp)
		}
		return err
	}
	return nil
}

//type AccountContext struct {
//	context.Context
//	AccountName string
//	UserId      repos.ID
//}

type UserContext struct {
	context.Context
	UserId repos.ID
}

//func GetUser(ctx AccountContext) (string, error) {
//	session := httpServer.GetSession[*common.AuthSession](ctx)
//	if session == nil {
//		return "", errors.New("Unauthorized")
//	}
//	return string(session.UserId), nil
//}

func (d *domain) checkAccountAccess(ctx context.Context, accountName string, userId repos.ID, action iamT.Action) error {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId:       string(userId),
		ResourceRefs: []string{iamT.NewResourceRef(accountName, iamT.ResourceAccount, accountName)},
		Action:       string(action),
	})

	if err != nil {
		return err
	}

	if !co.Status {
		return fmt.Errorf("unauthorized to access this account")
	}

	return nil
}
