package types

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/repos"
)

type InfraContext struct {
	context.Context
	Logger logging.Logger

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

const (
	GlobalVPNConnectionDeviceMethod = "gvpn-connection"
	KloudliteGlobalVPNDeviceMethod  = "kloudlite-global-vpn-device"
)

type ErrGRPCCall struct {
	Err error
}

func (e ErrGRPCCall) Error() string {
	return fmt.Sprintf("grpc call failed with error: %v", errors.NewE(e.Err))
}

type ErrIAMUnauthorized struct {
	UserId   string
	Resource string
	Action   string
}

func (e ErrIAMUnauthorized) Error() string {
	return fmt.Sprintf("user (%q) is unauthorized to perform action (%q) on resource (%q)", e.UserId, e.Action, e.Resource)
}
