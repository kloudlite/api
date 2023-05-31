package domain

import (
	mongogridfs "kloudlite.io/pkg/mongo-gridfs"
)

type CommonProviderData struct {
	// StorePath   string            `yaml:"storePath"`
	TfTemplates string            `yaml:"tfTemplates"`
	Labels      map[string]string `yaml:"labels"`
	Taints      []string          `yaml:"taints"`
	// Secrets     string            `yaml:"secrets"`
	SSHPath string `yaml:"sshPath"`
}

type AwsProviderConfig struct {
	AccessKey    string `yaml:"accessKey"`
	AccessSecret string `yaml:"accessSecret"`
	AccountId    string `yaml:"accountId"`
}

type AWSNode struct {
	NodeId       string `yaml:"nodeId"`
	Region       string `yaml:"region"`
	InstanceType string `yaml:"instanceType"`
	VPC          string `yaml:"vpc"`
	ImageId      string `yaml:"imageId"`
}

type awsClient struct {
	gfs  mongogridfs.GridFs
	node AWSNode

	accessKey    string
	accessSecret string

	SSHPath     string
	accountId   string
	secrets     string
	providerDir string
	storePath   string
	tfTemplates string
	labels      map[string]string
	taints      []string
}

// // getFolder implements doProviderClient
// func (a awsClient) getFolder() string {
// 	// eg -> /path/acc_id/do/blr1/node_id/do
//
// 	return path.Join(a.storePath, a.accountId, a.providerDir, a.node.Region, a.node.NodeId)
// }

// // initTFdir implements doProviderClient
// func (d awsClient) initTFdir() error {
//
// 	folder := d.getFolder()
//
// 	if err := utils.ExecCmd(fmt.Sprintf("cp -r %s %s", fmt.Sprintf("%s/%s", d.tfTemplates, d.providerDir), folder), "initialize terraform"); err != nil {
// 		return err
// 	}
//
// 	cmd := exec.Command("terraform", "init")
// 	cmd.Dir = path.Join(folder, d.providerDir)
//
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
//
// 	return cmd.Run()
// }

// NewNode implements ProviderClient
func (a awsClient) NewNode() error {

	values := map[string]string{}

	values["access_key"] = a.accessKey
	values["secret_key"] = a.accessSecret

	values["region"] = a.node.Region
	values["node_id"] = a.node.NodeId
	values["instance_type"] = a.node.InstanceType
	values["keys-path"] = a.SSHPath
	values["ami"] = a.node.ImageId

	/*
		steps:
			- check if state present in db
			- if present load that to working dir
			- else initialize new tf dir
			- apply terraform
			- upload the final state with defer
	*/

	// making dir
	// if err := utils.Mkdir(a.getFolder()); err != nil {
	// 	return err
	// }
	//
	// // initialize directory
	// if err := a.initTFdir(); err != nil {
	// 	return err
	// }
	//
	// // apply terraform
	// return utils.ApplyTF(path.Join(a.getFolder(), a.providerDir), values)
	//
	return nil
}

// AttachNode implements ProviderClientkkk
func (awsClient) AttachNode() error {
	panic("unimplemented")
}

// DeleteNode implements ProviderClient
func (awsClient) DeleteNode() error {
	panic("unimplemented")
}

// UnattachNode implements ProviderClient
func (awsClient) UnattachNode() error {
	panic("unimplemented")
}

func NewAwsProviderClient(node AWSNode, cpd CommonProviderData, apc AwsProviderConfig) ProviderClient {
	return awsClient{
		node:         node,
		accessKey:    apc.AccessKey,
		accessSecret: apc.AccessSecret,
		accountId:    apc.AccountId,

		providerDir: "aws",
		// secrets:     cpd.Secrets,
		// storePath:   cpd.StorePath,
		tfTemplates: cpd.TfTemplates,
		labels:      cpd.Labels,
		taints:      cpd.Taints,
		SSHPath:     cpd.SSHPath,
	}
}
