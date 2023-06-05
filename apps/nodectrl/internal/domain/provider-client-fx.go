package domain

import (
	"fmt"
	"io/ioutil"
	"os"

	"go.uber.org/fx"

	"kloudlite.io/apps/nodectrl/internal/domain/aws"
	"kloudlite.io/apps/nodectrl/internal/domain/common"
	"kloudlite.io/apps/nodectrl/internal/domain/do"
	"kloudlite.io/apps/nodectrl/internal/domain/utils"
	"kloudlite.io/apps/nodectrl/internal/env"
)

var ProviderClientFx = fx.Module("provider-client-fx",
	fx.Provide(func(env *env.Env, d Domain) (common.ProviderClient, error) {
		privateKeyBytes, publicKeyBytes, err := utils.GenerateKeys()
		if err != nil {
			return nil, err
		}

		const sshDir = "/tmp/ssh"

		if _, err := os.Stat(sshDir); err != nil {
			if e := os.Mkdir(sshDir, os.ModePerm); e != nil {
				return nil, e
			}
		}

		file, err := ioutil.TempDir("/tmp/ssh", "ssh_")
		if err != nil {
			return nil, err
		}

		if err := os.WriteFile(fmt.Sprintf("%s/access.pub", file), publicKeyBytes, os.ModePerm); err != nil {
			return nil, err
		}

		if err := os.WriteFile(fmt.Sprintf("%s/access", file), privateKeyBytes, os.ModePerm); err != nil {
			return nil, err
		}

		cpd := common.CommonProviderData{}

		if err := utils.Base64YamlDecode(env.ProviderConfig, &cpd); err != nil {
			return nil, err
		}

		switch env.CloudProvider {
		case "aws":

			node := aws.AWSNode{}

			if err := utils.Base64YamlDecode(env.NodeConfig, &node); err != nil {
				return nil, err
			}

			apc := aws.AwsProviderConfig{}

			fmt.Println("here......................", env.AWSProviderConfig)

			if err := utils.Base64YamlDecode(env.AWSProviderConfig, &apc); err != nil {
				return nil, err
			}

			return aws.NewAwsProviderClient(node, cpd, apc, d.GetGRidFs(), d.GetTokenRepo()), nil
		case "azure":
			panic("not implemented")
		case "do":

			node := do.DoNode{}

			if err := utils.Base64YamlDecode(env.NodeConfig, &node); err != nil {
				return nil, err
			}

			dpc := do.DoProviderConfig{}

			if err := utils.Base64YamlDecode(env.DoProviderConfig, &dpc); err != nil {
				return nil, err
			}

			return do.NewDoProviderClient(node, cpd, dpc, d.GetGRidFs(), d.GetTokenRepo()), nil
		case "gcp":
			panic("not implemented")
		}

		return nil, nil
	}),
)
