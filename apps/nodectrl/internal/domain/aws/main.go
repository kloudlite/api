package aws

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"kloudlite.io/apps/nodectrl/internal/domain/common"
	"kloudlite.io/apps/nodectrl/internal/domain/entities"
	"kloudlite.io/apps/nodectrl/internal/domain/utils"
	mongogridfs "kloudlite.io/pkg/mongo-gridfs"
	"kloudlite.io/pkg/repos"
)

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
	gfs       mongogridfs.GridFs
	node      AWSNode
	tokenRepo repos.DbRepo[*entities.Token]

	accessKey    string
	accessSecret string
	accountId    string

	SSHPath     string
	tfTemplates string
	labels      map[string]string
	taints      []string
}

// CreateAndAttachNode implements common.ProviderClient
func (a awsClient) CreateAndAttachNode(ctx context.Context) error {
	if err := a.NewNode(ctx); err != nil {
		return err
	}

	if err := a.AttachNode(ctx); err != nil {
		return err
	}

	return nil
}

// AttachNode implements common.ProviderClient
func (a awsClient) AttachNode(ctx context.Context) error {
	/*
		check readyness, wait if not ready
		if ready install agent
		  to install fetch
	*/

	token, err := a.tokenRepo.FindOne(ctx, repos.Filter{"nodeId": a.node.NodeId, "accountId": a.accountId})
	if err != nil {
		return err
	}

	var out []byte

	if out, err = utils.GetOutput(path.Join(utils.Workdir, a.node.NodeId), "node-ip"); err != nil {
		return err
	}

	labels := func() []string {
		l := []string{}
		for k, v := range a.labels {
			l = append(l, fmt.Sprintf("--node-label %s=%s", k, v))
		}
		l = append(l, fmt.Sprintf("--node-label %s=%s", "kloudlite.io/public-ip", string(out)))
		return l
	}()

	count := 0

	for {
		if e := utils.ExecCmd(
			fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s ls",
				fmt.Sprintf("%v/access", a.SSHPath),
				string(out)),
			"checking if node is ready "); e == nil {
			break
		}

		count++
		if count > 24 {
			return fmt.Errorf("node is not ready even after 6 minutes")
		}
		time.Sleep(time.Second * 15)
	}

	// attach node
	if e := utils.ExecCmd(
		fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s sudo sh /tmp/k3s-install.sh agent --server %s --token %s %s --node-name %s --node-external-ip %s --node-ip %s",
			fmt.Sprintf("%v/access", a.SSHPath), string(out), token.EndpointUrl, token.JoinToken,
			strings.Join(labels, " "), a.node.NodeId, string(out), string(out)),
		"attaching to cluster"); e != nil {
		return e
	}

	return nil
}

// CreateCluster implements common.ProviderClient
func (a awsClient) CreateCluster(ctx context.Context) error {
	/*
		create node
		check for rediness
		install k3s
		check for rediness
		install maaster
	*/

	const sshDir = "/tmp/ssh"

	if _, err := os.Stat(sshDir); err != nil {
		return os.Mkdir(sshDir, os.ModePerm)
	}

	file, err := ioutil.TempDir("/tmp/ssh", "ssh_")
	if err != nil {
		return err
	}

	if err := a.NewNode(ctx); err != nil {
		return err
	}

	var ip []byte

	if ip, err = utils.GetOutput(path.Join(utils.Workdir, a.node.NodeId), "node-ip"); err != nil {
		return err
	}

	count := 0

	for {
		if e := utils.ExecCmd(
			fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s ls",
				fmt.Sprintf("%v/access", file),
				string(ip)),
			"checking if node is ready "); e == nil {
			break
		}

		count++
		if count > 24 {
			return fmt.Errorf("node is not ready even after 6 minutes")
		}
		time.Sleep(time.Second * 15)
	}

	// install k3s
	cmd := fmt.Sprintf(
		"ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s sudo sh /tmp/k3s-install.sh server --token=%q --node-external-ip %s --flannel-backend wireguard-native --flannel-external-ip --disable traefik --node-name=%q",
		file,
		string(ip),
		a.node.NodeId,
		string(ip),
		fmt.Sprintf("kl-master-%s", a.node.NodeId),
	)

	if err := utils.ExecCmd(cmd, cmd); err != nil {
		return err
	}
	// needed to fetch kubeconfig

	configOut, err := utils.ExecCmdWithOutput(fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s cat /etc/rancher/k3s/k3s.yaml", file, string(ip)), "")
	if err != nil {
		return err
	}

	var kubeconfig common.KubeConfigType
	if err := yaml.Unmarshal(configOut, &kubeconfig); err != nil {
		return err
	}

	for i := range kubeconfig.Clusters {
		kubeconfig.Clusters[i].Cluster.Server = fmt.Sprintf("https://%s:6443", string(ip))
	}

	kc, err := yaml.Marshal(kubeconfig)
	if err != nil {
		return err
	}

	tokenOut, err := utils.ExecCmdWithOutput(fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s cat /var/lib/rancher/k3s/server/node-token", file, string(ip)), "")
	if err != nil {
		return err
	}

	_, err = a.tokenRepo.Create(ctx, &entities.Token{
		JoinToken:   string(tokenOut),
		EndpointUrl: fmt.Sprintf("https://%s:6443", ip),
		KubeConfig:  string(kc),
		NodeId:      a.node.NodeId,
		AccountName: a.accountId,
		ClusterName: "",
	})

	return err
}

func parseValues(a awsClient) map[string]string {
	values := map[string]string{}

	values["access_key"] = a.accessKey
	values["secret_key"] = a.accessSecret

	values["region"] = a.node.Region
	values["node_id"] = a.node.NodeId
	values["instance_type"] = a.node.InstanceType
	values["keys-path"] = a.SSHPath
	values["ami"] = a.node.ImageId

	fmt.Print(values)

	return values
}

func (a awsClient) SaveToDbGuranteed(ctx context.Context) {
	for {
		if err := utils.SaveToDb(ctx, a.node.NodeId, a.gfs); err == nil {
			break
		} else {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 20)
	}
}

// NewNode implements ProviderClient
func (a awsClient) NewNode(ctx context.Context) error {
	file, err := ioutil.TempDir("/tmp/ssh", "ssh_")
	if err != nil {
		return err
	}

	a.SSHPath = file

	values := parseValues(a)

	/*
		steps:
			- check if state present in db
			- if present load that to working dir
			- else initialize new tf dir
			- apply terraform
			- upload the final state with defer
	*/

	if err := utils.MakeTfWorkFileReady(ctx, a.node.NodeId, path.Join(a.tfTemplates, "aws"), a.gfs, true); err != nil {
		return err
	}

	defer a.SaveToDbGuranteed(ctx)

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
func (a awsClient) DeleteNode(ctx context.Context) error {
	values := parseValues(a)

	/*
		steps:
			- check if state present in db
			- if present load that to working dir
			- else initialize new tf dir
			- destroy node with terraform
			- delete final state
	*/

	if err := utils.MakeTfWorkFileReady(ctx, a.node.NodeId, path.Join(a.tfTemplates, "aws"), a.gfs, false); err != nil {
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

	filename := fmt.Sprintf("%s.zip", a.node.NodeId)

	if err := a.gfs.DeleteAllWithFilename(filename); err != nil {
		return err
	}

	return nil
}

func NewAwsProviderClient(node AWSNode, cpd common.CommonProviderData, apc AwsProviderConfig, gfs mongogridfs.GridFs, tokenRepo repos.DbRepo[*entities.Token]) common.ProviderClient {
	return awsClient{
		node:      node,
		gfs:       gfs,
		tokenRepo: tokenRepo,

		accessKey:    apc.AccessKey,
		accessSecret: apc.AccessSecret,
		accountId:    apc.AccountId,

		tfTemplates: cpd.TfTemplates,
		labels:      cpd.Labels,
		taints:      cpd.Taints,
		SSHPath:     cpd.SSHPath,
	}
}
