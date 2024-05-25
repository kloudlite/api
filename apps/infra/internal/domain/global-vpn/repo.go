package global_vpn

import (
	"github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fc "github.com/kloudlite/api/apps/infra/internal/entities/field-constants"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

type IncrementAllocatedGlobalVPNDevicesArgs struct {
	IncrementBy   int
	GlobalVPNId   repos.ID
	GlobalVPNRepo repos.DbRepo[*entities.GlobalVPN]
}

func IncrementAllocatedGlobalVPNDevices(ctx types.InfraContext, args IncrementAllocatedGlobalVPNDevicesArgs) error {
	if _, err := args.GlobalVPNRepo.PatchById(ctx, args.GlobalVPNId, repos.Document{"$inc": map[string]any{fc.GlobalVPNNumAllocatedDevices: 1}}); err != nil {
		return errors.NewE(err)
	}

	return nil
}

type GetGlobalVPNArgs struct {
	GlobalVPNName string
	GlobalVPNRepo repos.DbRepo[*entities.GlobalVPN]
}

func GetGlobalVPN(ctx types.InfraContext, args GetGlobalVPNArgs) (*entities.GlobalVPN, error) {
	return args.GlobalVPNRepo.FindOne(ctx, entities.UniqueGlobalVPN(ctx.AccountName, args.GlobalVPNName))
}
