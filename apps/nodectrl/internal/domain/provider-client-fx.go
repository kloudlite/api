package domain

import (
	"go.uber.org/fx"
	"kloudlite.io/apps/nodectrl/internal/domain/utils"
	"kloudlite.io/apps/nodectrl/internal/env"
)

var ProviderClientFx = fx.Module("provider-client-fx",
	fx.Provide(func(env *env.Env) (ProviderClient, error) {

		cpd := CommonProviderData{}

		if err := utils.Base64YamlDecode(env.ProviderConfig, &cpd); err != nil {
			return nil, err
		}

		switch env.CloudProvider {
		case "aws":

			node := AWSNode{}

			if err := utils.Base64YamlDecode(env.NodeConfig, &node); err != nil {
				return nil, err
			}

			apc := AwsProviderConfig{}

			if err := utils.Base64YamlDecode(env.AWSProviderConfig, &apc); err != nil {
				return nil, err
			}

			return NewAwsProviderClient(node, cpd, apc), nil
		case "azure":
			panic("not implemented")
		case "do":
			panic("not implemented")
		case "gcp":
			panic("not implemented")
		}
		return awsClient{}, nil
	}),
)
