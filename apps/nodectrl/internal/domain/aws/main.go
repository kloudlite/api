package aws

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"kloudlite.io/apps/nodectrl/internal/domain/common"
	"kloudlite.io/apps/nodectrl/internal/domain/utils"
	awss3 "kloudlite.io/pkg/aws-s3"
)

type AwsProviderConfig struct {
	AccessKey    string `yaml:"accessKey"`
	AccessSecret string `yaml:"accessSecret"`
	AccountName  string `yaml:"accountName"`
}

type AWSNode struct {
	NodeId       string `yaml:"nodeId"`
	Region       string `yaml:"region"`
	InstanceType string `yaml:"instanceType"`
	VPC          string `yaml:"vpc"`
	ImageId      string `yaml:"imageId"`
	IsGpu        bool   `yaml:"isGpu"`
	NodeType     string `yaml:"nodeType" json:"nodeType"`
}

type AwsClient struct {
	node        AWSNode
	awsS3Client awss3.AwsS3

	accessKey    string
	accessSecret string
	accountName  string

	// SSHPath     string
	tfTemplates string
	labels      map[string]string
	taints      []string
}

type TokenAndKubeconfig struct {
	Token       string `json:"token"`
	Kubeconfig  string `json:"kubeconfig"`
	ServerIp    string `json:"serverIp"`
	MasterToken string `json:"masterToken"`
}

func (a AwsClient) SetupSSH() error {
	const sshDir = "/tmp/ssh"

	if _, err := os.Stat(sshDir); err != nil {
		return os.Mkdir(sshDir, os.ModePerm)
	}

	destDir := path.Join(sshDir, a.accountName)
	fileName := fmt.Sprintf("%s.zip", a.accountName)

	if err := a.awsS3Client.IsFileExists(fileName); err != nil {

		if _, err := os.Stat(destDir); err == nil {
			if err := os.RemoveAll(destDir); err != nil {
				return err
			}
		}

		if e := os.Mkdir(destDir, os.ModePerm); e != nil {
			return e
		}

		privateKeyBytes, publicKeyBytes, err := utils.GenerateKeys()
		if err != nil {
			return err
		}

		if err := os.WriteFile(fmt.Sprintf("%s/access.pub", destDir), publicKeyBytes, os.ModePerm); err != nil {
			return err
		}

		if err := os.WriteFile(fmt.Sprintf("%s/access", destDir), privateKeyBytes, 0400); err != nil {
			return err
		}
		return nil
	}

	if err := os.RemoveAll(destDir); err != nil {
		return err
	}

	err := a.awsS3Client.DownloadFile(path.Join(sshDir, fileName), fileName)
	if err != nil {
		return err
	}

	_, err = utils.Unzip(path.Join(sshDir, fileName), sshDir)
	if err != nil {
		return err
	}

	return nil
}

func (a AwsClient) saveForSure() error {
	count := 0
	for {
		if err := a.saveSSH(); err == nil {
			return nil
		}
		if count >= 10 {
			return fmt.Errorf("coudn't save the state")
		}

		time.Sleep(time.Second * 20)
		count++
	}
}

func (a AwsClient) saveSSH() error {
	const sshDir = "/tmp/ssh"
	destDir := path.Join(sshDir, a.accountName)
	fileName := fmt.Sprintf("%s.zip", a.accountName)

	if err := utils.ZipSource(destDir, path.Join(sshDir, fileName)); err != nil {
		return err
	}

	if err := a.awsS3Client.UploadFile(path.Join(sshDir, fileName), fileName); err != nil {
		return err
	}

	return nil
}

func (a AwsClient) SaveToDbGuranteed(ctx context.Context) {
	for {
		if err := utils.SaveToDb(a.node.NodeId, a.awsS3Client); err == nil {
			break
		} else {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 20)
	}
}

// NewNode implements ProviderClient
func (a AwsClient) NewNode(ctx context.Context) error {
	sshPath := path.Join("/tmp/ssh", a.accountName)
	values := parseValues(a, sshPath)

	if true {
		if err := utils.MakeTfWorkFileReady(a.node.NodeId, path.Join(a.tfTemplates, "aws"), a.awsS3Client, true); err != nil {
			return err
		}

		defer a.SaveToDbGuranteed(ctx)
	}

	// upload the final state to the db, upsert if db is already present

	// apply the tf file
	if err := func() error {
		if err := utils.InitTFdir(path.Join(utils.Workdir, a.node.NodeId)); err != nil {
			return err
		}

		if err := utils.ApplyTF(path.Join(utils.Workdir, a.node.NodeId), values); err != nil {
			return err
		}

		return nil
	}(); err != nil {
		return err
	}

	return nil
}

// DeleteNode implements ProviderClient
func (a AwsClient) DeleteNode(ctx context.Context) error {
	sshPath := path.Join("/tmp/ssh", a.accountName)
	values := parseValues(a, sshPath)

	/*
		steps:
			- check if state present in db
			- if present load that to working dir
			- else initialize new tf dir
			- destroy node with terraform
			- delete final state
	*/

	if err := utils.MakeTfWorkFileReady(a.node.NodeId, path.Join(a.tfTemplates, "aws"), a.awsS3Client, false); err != nil {
		return err
	}

	// destroy the tf file
	if err := func() error {
		if err := utils.DestroyNode(a.node.NodeId, values); err != nil {
			return err
		}

		return nil
	}(); err != nil {
		return err
	}

	return nil
}

func NewAwsProviderClient(node AWSNode, cpd common.CommonProviderData, apc AwsProviderConfig) (common.ProviderClient, error) {
	awsS3Client, err := awss3.NewAwsS3Client(apc.AccessKey, apc.AccessSecret, apc.AccountName)
	if err != nil {
		return nil, err
	}

	return AwsClient{
		node:        node,
		awsS3Client: awsS3Client,

		accessKey:    apc.AccessKey,
		accessSecret: apc.AccessSecret,
		accountName:  apc.AccountName,

		tfTemplates: cpd.TfTemplates,
		labels:      cpd.Labels,
		taints:      cpd.Taints,
	}, nil
}
