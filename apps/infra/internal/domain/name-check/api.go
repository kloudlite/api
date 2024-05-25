package name_check

import (
	"context"

	domainT "github.com/kloudlite/api/apps/infra/internal/domain/types"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
)

func (d *Domain) SuggestName(ctx context.Context, seed *string) string {
	return fn.GenReadableName(fn.DefaultIfNil(seed, ""))
}

type ResType string

const (
	ResTypeCluster               ResType = "cluster"
	ResTypeBYOKCluster           ResType = "byok_cluster"
	ResTypeGlobalVPNDevice       ResType = "global_vpn_device"
	ResTypeClusterManagedService ResType = "cluster_managed_service"
	ResTypeProviderSecret        ResType = "providersecret"
	ResTypeNodePool              ResType = "nodepool"
	ResTypeHelmRelease           ResType = "helm_release"
)

type CheckNameAvailabilityOutput struct {
	Result         bool     `json:"result"`
	SuggestedNames []string `json:"suggestedNames"`
}

func (d *Domain) CheckNameAvailability(ctx domainT.InfraContext, typeArg ResType, clusterName *string, name string) (*CheckNameAvailabilityOutput, error) {
	if !fn.IsValidK8sResourceName(name) {
		return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
	}

	switch typeArg {
	case ResTypeCluster:
		{
			ok, err := d.IsClusterNameAvailable(ctx, name)
			if err != nil {
				return nil, errors.NewE(err)
			}
			if ok {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}
			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
		}
	case ResTypeBYOKCluster:
		{
			ok, err := d.IsClusterNameAvailable(ctx, name)
			if err != nil {
				return nil, errors.NewE(err)
			}
			if ok {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}
			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
		}
	case ResTypeProviderSecret:
		{
			ok, err := d.IsProviderSecretNameAvailable(ctx, name)
			if err != nil {
				return nil, errors.NewE(err)
			}
			if ok {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}
			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
		}
	case ResTypeGlobalVPNDevice:
		{
			ok, err := d.IsGlobalVPNDeviceNameAvailable(ctx, name)
			if err != nil {
				return nil, errors.NewE(err)
			}
			if ok {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}
			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
		}
	case ResTypeNodePool:
		{
			if clusterName == nil || *clusterName == "" {
				return nil, errors.Newf("clusterName is required for checking name availability for %s", ResTypeHelmRelease)
			}

			ok, err := d.IsNodepoolNameAvailable(ctx, *clusterName, name)
			if err != nil {
				return nil, errors.NewE(err)
			}
			if ok {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}
			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
		}
	case ResTypeHelmRelease:
		{
			if clusterName == nil || *clusterName == "" {
				return nil, errors.Newf("clusterName is required for checking name availability for %s", ResTypeNodePool)
			}

			ok, err := d.IsHelmReleaseNameAvailable(ctx, *clusterName, name)
			if err != nil {
				return nil, errors.NewE(err)
			}
			if ok {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}
			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
		}
	case ResTypeClusterManagedService:
		{
			if clusterName == nil || *clusterName == "" {
				return nil, errors.Newf("clusterName is required for checking name availability for %s", ResTypeClusterManagedService)
			}

			ok, err := d.IsClusterManagedSvcNameAvailable(ctx, *clusterName, name)
			if err != nil {
				return nil, errors.NewE(err)
			}
			if ok {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}
			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: fn.GenValidK8sResourceNames(name, 3)}, nil
		}
	default:
		{
			return &CheckNameAvailabilityOutput{Result: false}, errors.Newf("unknown resource type provided: %q", typeArg)
		}
	}
}
