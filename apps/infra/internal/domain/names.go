package domain

import (
	"context"
	"fmt"

	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/repos"
)

func (d *domain) SuggestName(ctx context.Context, seed *string) string {
	return fn.GenReadableName(fn.DefaultIfNil(seed, ""))
}

type ResType string

const (
	ResTypeCluster ResType = "cluster"
	// ResTypeCloudProvider  ResType = "cloudprovider"
	// ResTypeEdge           ResType = "edge"
	ResTypeProviderSecret ResType = "providersecret"
	ResTypeNodePool       ResType = "nodepool"
)

type CheckNameAvailabilityOutput struct {
	Result         bool     `json:"result"`
	SuggestedNames []string `json:"suggestedNames"`
}

func (d *domain) CheckNameAvailability(ctx InfraContext, typeArg ResType, name string) (*CheckNameAvailabilityOutput, error) {
	switch typeArg {
	case ResTypeCluster:
		{
			cp, err := d.clusterRepo.FindOne(ctx, repos.Filter{
				"accountName":   ctx.AccountName,
				"metadata.name": name,
			})
			if err != nil {
				return &CheckNameAvailabilityOutput{Result: false}, err
			}

			if cp == nil {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}

			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: []string{
				fn.GenReadableName(name), fn.GenReadableName(name), fn.GenReadableName(name),
			}}, nil
		}
	case ResTypeProviderSecret:
		{
			cp, err := d.secretRepo.FindOne(ctx, repos.Filter{
				"accountName":   ctx.AccountName,
				"metadata.name": name,
			})
			if err != nil {
				return &CheckNameAvailabilityOutput{Result: false}, err
			}

			if cp == nil {
				return &CheckNameAvailabilityOutput{Result: true}, nil
			}

			return &CheckNameAvailabilityOutput{Result: false, SuggestedNames: []string{
				fn.GenReadableName(name), fn.GenReadableName(name), fn.GenReadableName(name),
			}}, nil
		}
	case ResTypeNodePool:
		{
			return nil, fmt.Errorf("not implemented")
		}
	default:
		{
			return &CheckNameAvailabilityOutput{Result: false}, nil
		}
	}
}
