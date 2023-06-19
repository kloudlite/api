package aws

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	guuid "github.com/google/uuid"
	"gopkg.in/yaml.v2"

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

// AddMaster implements common.ProviderClient.
func (a AwsClient) AddMaster(ctx context.Context) error {
	// fetch token
	sshPath := path.Join("/tmp/ssh", a.accountName)

	tokenFileName := fmt.Sprintf("%s-config.yaml", a.accountName)

	if err := a.awsS3Client.IsFileExists(tokenFileName); err != nil {
		return err
	}

	if _, err := os.Stat(sshPath); err != nil {
		if e := os.Mkdir(sshPath, os.ModePerm); e != nil {
			return e
		}
	}

	tokenPath := path.Join(sshPath, "config.yaml")
	if err := a.awsS3Client.DownloadFile(tokenPath, tokenFileName); err != nil {
		return err
	}

	b, err := os.ReadFile(tokenPath)
	if err != nil {
		return err
	}

	kc := TokenAndKubeconfig{}

	if err := yaml.Unmarshal(b, &kc); err != nil {
		return err
	}

	// setup ssh

	if err := a.SetupSSH(); err != nil {
		return err
	}
	defer a.saveForSure()

	// create node and wait for ready
	if err := a.NewNode(ctx); err != nil {
		return err
	}

	ip, err := utils.GetOutput(path.Join(utils.Workdir, a.node.NodeId), "node-ip")
	if err != nil {
		return err
	}

	count := 0

	for {
		if e := utils.ExecCmd(
			fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s ls",
				fmt.Sprintf("%v/access", sshPath),
				string(ip),
			),
			"checking if node is ready"); e == nil {
			break
		}

		count++
		if count > 24 {
			return fmt.Errorf("node is not ready even after 6 minutes")
		}
		time.Sleep(time.Second * 5)
	}

	// attach to cluster as master
	cmd := fmt.Sprintf(
		"ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s/access root@%s sudo sh /tmp/k3s-install.sh server --server https://%s:6443 --token %s  --node-external-ip %s --flannel-backend wireguard-native --flannel-external-ip --disable traefik --node-name=%s",
		sshPath,
		string(ip),
		kc.ServerIp,
		strings.TrimSpace(string(kc.Token)),
		string(ip),
		a.node.NodeId,
	)

	if err := utils.ExecCmd(cmd, "attaching to cluster as a master"); err != nil {
		return err
	}

	return nil
}

func (a AwsClient) AddWorker(ctx context.Context) error {
	// fetch token

	sshPath := path.Join("/tmp/ssh", a.accountName)

	if _, err := os.Stat(sshPath); err != nil {
		if e := os.Mkdir(sshPath, os.ModePerm); e != nil {
			return e
		}
	}

	tokenFileName := fmt.Sprintf("%s-config.yaml", a.accountName)

	if err := a.awsS3Client.IsFileExists(tokenFileName); err != nil {
		return err
	}

	tokenPath := path.Join(sshPath, "config.yaml")
	if err := a.awsS3Client.DownloadFile(tokenPath, tokenFileName); err != nil {
		return err
	}

	b, err := os.ReadFile(tokenPath)
	if err != nil {
		return err
	}

	kc := TokenAndKubeconfig{}

	if err := yaml.Unmarshal(b, &kc); err != nil {
		return err
	}

	// setup ssh

	if err := a.SetupSSH(); err != nil {
		return err
	}
	defer a.saveForSure()

	// create node and wait for ready
	if err := a.NewNode(ctx); err != nil {
		return err
	}

	ip, err := utils.GetOutput(path.Join(utils.Workdir, a.node.NodeId), "node-ip")
	if err != nil {
		return err
	}

	count := 0

	for {
		if e := utils.ExecCmd(
			fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s ls",
				fmt.Sprintf("%s/access", sshPath),
				string(ip),
			),
			"checking if node ready"); e == nil {
			break
		}

		count++
		if count > 24 {
			return fmt.Errorf("node is not ready even after 6 minutes")
		}
		time.Sleep(time.Second * 5)
	}

	labels := func() []string {
		l := []string{}
		for k, v := range map[string]string{
			"kloudlite.io/public-ip": string(ip),
		} {
			l = append(l, fmt.Sprintf("--node-label %s=%s", k, v))
		}

		for k, v := range a.labels {
			l = append(l, fmt.Sprintf("--node-label %s=%s", k, v))
		}
		return l
	}()

	// attach to cluster as workernode

	cmd := fmt.Sprintf(
		"ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s/access root@%s sudo sh /tmp/k3s-install.sh agent --server https://%s:6443 --token=%s --node-external-ip %s --node-name %s %s %s",
		sshPath,
		ip,
		kc.ServerIp,
		strings.TrimSpace(string(kc.Token)),
		ip,
		a.node.NodeId,
		strings.Join(labels, " "),
		func() string {
			if a.node.IsGpu {
				// return "--docker"
				// return "--docker"
				return ""
			}
			return ""
		}(),
	)

	if err := utils.ExecCmd(cmd, "attaching to cluster as a worker node"); err != nil {
		return err
	}

	return nil
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

// CreateCluster implements common.ProviderClient
func (a AwsClient) CreateCluster(ctx context.Context) error {
	/*
		create node
		check for rediness
		install k3s
		check for rediness
		install maaster
	*/

	if err := a.SetupSSH(); err != nil {
		return err
	}
	defer a.saveForSure()
	sshPath := path.Join("/tmp/ssh", a.accountName)

	if err := a.NewNode(ctx); err != nil {
		return err
	}

	ip, err := utils.GetOutput(path.Join(utils.Workdir, a.node.NodeId), "node-ip")
	if err != nil {
		return err
	}

	count := 0

	for {
		if e := utils.ExecCmd(
			fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s root@%s ls",
				fmt.Sprintf("%v/access", sshPath),
				string(ip),
			),
			"checking is node is ready"); e == nil {
			break
		}

		count++
		if count > 24 {
			return fmt.Errorf("node is not ready even after 6 minutes")
		}
		time.Sleep(time.Second * 5)
	}

	masterToken := guuid.New()

	// install k3s
	cmd := fmt.Sprintf(
		"ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s/access root@%s sudo sh /tmp/k3s-install.sh server --token=%s --node-external-ip %s --flannel-backend wireguard-native --flannel-external-ip --disable traefik --node-name=%s --cluster-init",
		sshPath,
		string(ip),
		masterToken.String(),
		string(ip),
		a.node.NodeId,
	)

	if err := utils.ExecCmd(cmd, "installing k3s"); err != nil {
		return err
	}
	// needed to fetch kubeconfig

	configOut, err := utils.ExecCmdWithOutput(fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s/access root@%s cat /etc/rancher/k3s/k3s.yaml", sshPath, string(ip)), "fetching kubeconfig from the cluster")
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

	tokenOut, err := utils.ExecCmdWithOutput(fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i %s/access root@%s cat /var/lib/rancher/k3s/server/node-token", sshPath, string(ip)), "fetching node token from the cluster")
	if err != nil {
		return err
	}

	st := TokenAndKubeconfig{
		Token:       string(tokenOut),
		Kubeconfig:  string(kc),
		ServerIp:    string(ip),
		MasterToken: masterToken.String(),
	}

	b, err := yaml.Marshal(st)
	if err != nil {
		return err
	}

	tokenPath := path.Join(sshPath, "config.yaml")

	if err := os.WriteFile(tokenPath, b, os.ModePerm); err != nil {
		return err
	}

	if err := a.awsS3Client.UploadFile(tokenPath, fmt.Sprintf("%s-config.yaml", a.accountName)); err != nil {
		return err
	}

	return err
}

func parseValues(a AwsClient, sshPath string) map[string]string {
	values := map[string]string{}

	values["access_key"] = a.accessKey
	values["secret_key"] = a.accessSecret

	values["region"] = a.node.Region
	values["node_id"] = a.node.NodeId
	values["instance_type"] = a.node.InstanceType
	values["keys-path"] = sshPath
	values["ami"] = a.node.ImageId

	return values
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
