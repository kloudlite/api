package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/kloudlite/operator/pkg/constants"
	"sigs.k8s.io/yaml"
)

type AgentConfig struct {
	PublicIP string            `json:"publicIP"`
	ServerIP string            `json:"serverIP"`
	Token    string            `json:"token"`
	NodeName string            `json:"nodeName"`
	Taints   []string          `json:"taints"`
	Labels   map[string]string `json:"labels"`
}

type PrimaryMasterConfig struct {
	PublicIP string            `json:"publicIP"`
	Token    string            `json:"token"`
	NodeName string            `json:"nodeName"`
	Taints   []string          `json:"taints"`
	Labels   map[string]string `json:"labels"`
	SANs     []string          `json:"SANs"`
}

type SecondaryMasterConfig struct {
	PublicIP string            `json:"publicIP"`
	ServerIP string            `json:"serverIP"`
	Token    string            `json:"token"`
	NodeName string            `json:"nodeName"`
	Taints   []string          `json:"taints"`
	Labels   map[string]string `json:"labels"`
	SANs     []string          `json:"SANs"`
}

type RunAsMode string

const (
	RunAsAgent           RunAsMode = "agent"
	RunAsPrimaryMaster   RunAsMode = "primaryMaster"
	RunAsSecondaryMaster RunAsMode = "secondaryMaster"
)

type K3sRunnerConfig struct {
	RunAs           RunAsMode              `json:"runAs"`
	Agent           *AgentConfig           `json:"agent"`
	PrimaryMaster   *PrimaryMasterConfig   `json:"primaryMaster"`
	SecondaryMaster *SecondaryMasterConfig `json:"secondaryMaster"`
}

func ColorText(text interface{}, code int) string {
	return fmt.Sprintf("\033[38;05;%dm%v\033[0m", code, text)
}

func execK3s(ctx context.Context, args ...string) error {
	stdout, err := os.Create("runner.stdout.log")
	if err != nil {
		return err
	}

	stderr, err := os.Create("runner.stderr.log")
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "k3s", args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	fmt.Printf("executing this shell command: %s\n", cmd.String())
	if err := cmd.Run(); err != nil {
		fmt.Printf("[ERROR]: %s", err.Error())
		return err
	}
	return nil
}

func ExecCmd2(cmdString string, logStr string) error {
	r := csv.NewReader(strings.NewReader(cmdString))
	r.Comma = ' '
	cmdArr, err := r.Read()
	if err != nil {
		return err
	}

	if logStr != "" {
		fmt.Printf("[#] %s\n", logStr)
	} else {
		fmt.Printf("[#] %s\n", strings.Join(cmdArr, " "))
	}

	cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
	cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("err occurred: %v\n", err.Error())
		return err
	}
	return nil
}

func getPublicIPv4() (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://ifconfig.me", nil)
	if err != nil {
		return "", err
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	return string(b), nil
}

func main() {
	runnerCfgFile := "/runner-config.yml"

	ctx, cf := signal.NotifyContext(context.TODO(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer cf()

	for {
		if err := ctx.Err(); err != nil {
			fmt.Println("context cancelled")
			os.Exit(1)
		}
		f, err := os.Open(runnerCfgFile)
		if err != nil {
			fmt.Println(ColorText(err.Error(), 1))
			continue
		}

		fmt.Println("found runner config file")

		out, err := io.ReadAll(f)
		if err != nil {
			fmt.Println(ColorText(err.Error(), 1))
			continue
		}

		var runnerCfg K3sRunnerConfig
		if err := yaml.Unmarshal(out, &runnerCfg); err != nil {
			fmt.Println(ColorText(err.Error(), 1))
			continue
		}

		switch runnerCfg.RunAs {
		case RunAsAgent:
			{
				if err := StartK3sAgent(ctx, runnerCfg.Agent); err != nil {
					if !errors.Is(err, context.Canceled) {
						fmt.Println(ColorText(err.Error(), 1))
						fmt.Println(ColorText("will retry after 10 second", 2))
						time.Sleep(time.Second * 10)
					}
				}
			}

		case RunAsPrimaryMaster:
			{
				if err := StartPrimaryK3sMaster(ctx, runnerCfg.PrimaryMaster); err != nil {
					if !errors.Is(err, context.Canceled) {
						fmt.Println(ColorText(err.Error(), 1))
						fmt.Println(ColorText("will retry after 10 second", 2))
						time.Sleep(time.Second * 10)
					}
				}
			}

		case RunAsSecondaryMaster:
			{
				if err := StartSecondaryK3sMaster(ctx, runnerCfg.SecondaryMaster); err != nil {
					if !errors.Is(err, context.Canceled) {
						fmt.Println(ColorText(err.Error(), 1))
						fmt.Println(ColorText("will retry after 10 second", 2))
						time.Sleep(time.Second * 10)
					}
				}
			}
		default:
			{
				fmt.Println(ColorText("invalid runAs mode", 1))
				continue
			}
		}

		fmt.Println(ColorText("Successfully Installed", 2))
		break
	}
}

func ExecCmdWithOutput(cmdString string, logStr string) ([]byte, error) {
	r := csv.NewReader(strings.NewReader(cmdString))
	r.Comma = ' '
	cmdArr, err := r.Read()
	if err != nil {
		return nil, err
	}

	if logStr != "" {
		fmt.Printf("[#] %s\n", logStr)
	} else {
		fmt.Printf("[#] %s\n", strings.Join(cmdArr, " "))
	}

	cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
	cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout

	return cmd.Output()
}

func StartPrimaryK3sMaster(ctx context.Context, pmc *PrimaryMasterConfig) error {
	fmt.Printf("starting as primary master, with configuration: %#v\n", *pmc)

	argsAndFlags := []string{
		"server",
		"--cluster-init",
		"--disable-helm-controller",
		"--disable", "traefik",
		"--disable", "servicelb",
		"--flannel-backend", "wireguard-native",
		"--node-name", pmc.NodeName,
		"--token", pmc.Token,
		"--write-kubeconfig-mode", "644",
		"--node-label", fmt.Sprintf("%s=%s", constants.PublicIpKey, pmc.PublicIP),
		"--node-label", fmt.Sprintf("%s=%s", constants.NodeNameKey, pmc.NodeName),
		"--tls-san", pmc.PublicIP,
	}

	for i := range pmc.SANs {
		argsAndFlags = append(argsAndFlags, "--tls-san", pmc.SANs[i])
	}

	for k, v := range pmc.Labels {
		argsAndFlags = append(argsAndFlags, "--node-label", fmt.Sprintf("%s=%s", k, v))
	}

	return execK3s(ctx, argsAndFlags...)
}

func StartSecondaryK3sMaster(ctx context.Context, smc *SecondaryMasterConfig) error {
	argsAndFlags := []string{
		"server",
		"--server", fmt.Sprintf("https://%s:6443", smc.ServerIP),
		"--disable-helm-controller",
		"--disable", "traefik",
		"--disable", "servicelb",
		"--flannel-backend", "wireguard-native",
		"--node-name", smc.NodeName,
		"--token", smc.Token,
		"--write-kubeconfig-mode", "644",
		"--node-label", fmt.Sprintf("%s=%s", constants.PublicIpKey, smc.PublicIP),
		"--node-label", fmt.Sprintf("%s=%s", constants.NodeNameKey, smc.NodeName),
		"--tls-san", smc.PublicIP,
	}

	for i := range smc.SANs {
		argsAndFlags = append(argsAndFlags, "--tls-san", smc.SANs[i])
	}

	for k, v := range smc.Labels {
		argsAndFlags = append(argsAndFlags, "--node-label", fmt.Sprintf("%s=%s", k, v))
	}

	return execK3s(ctx, argsAndFlags...)
}

func StartK3sAgent(ctx context.Context, agentCfg *AgentConfig) error {
	ip, err := func() (string, error) {
		if agentCfg.PublicIP != "" {
			return agentCfg.PublicIP, nil
		}

		return getPublicIPv4()
	}()

	if err != nil {
		return err
	}

	argsAndFlags := []string{
		"agent",
		"--server", fmt.Sprintf("https://%s:6443", agentCfg.ServerIP),
		"--token", agentCfg.Token,
		"--node-name", agentCfg.NodeName,
		"--node-label", fmt.Sprintf("%s=%s", constants.PublicIpKey, ip),
		"--node-label", fmt.Sprintf("%s=%s", constants.NodeNameKey, agentCfg.NodeName),
	}

	return execK3s(ctx, argsAndFlags...)
}
