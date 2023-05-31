package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"kloudlite.io/apps/nodectrl/internal/domain/utils"
)

const (
	CLUSTER_ID = "kl"
)

// rmTFdir implements doProviderClient
// func rmdir(folder string) error {
// 	return execCmd(fmt.Sprintf("rm -rf %q", folder), "")
// }

// makeTFdir implements doProviderClient
func mkdir(folder string) error {
	return utils.ExecCmd(fmt.Sprintf("mkdir -p %q", folder), "mkdir <terraform_dir>")
}

// destroyNode implements doProviderClient
func destroyNode(folder string, values map[string]string) error {
	vars := []string{"destroy", "-auto-approve"}

	for k, v := range values {
		vars = append(vars, fmt.Sprintf("-var=%s=%s", k, v))
	}

	cmd := exec.Command("terraform", vars...)
	cmd.Dir = folder

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return os.RemoveAll(folder)

	// return err
}

func getOutput(folder, key string) ([]byte, error) {
	vars := []string{"output", "-json"}
	fmt.Printf("[#] terraform %s\n", strings.Join(vars, " "))
	cmd := exec.Command("terraform", vars...)
	cmd.Dir = folder

	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return nil, err

	}

	// fmt.Println(string(out))

	var resp map[string]struct {
		Value string `json:"value"`
	}

	err = json.Unmarshal(out, &resp)
	if err != nil {
		return nil, err
	}

	return []byte(resp[key].Value), nil
}

// applyTF implements doProviderClient
func applyTF(folder string, values map[string]string) error {

	vars := []string{"apply", "-auto-approve"}

	fmt.Printf("[#] terraform %s", strings.Join(vars, " "))

	for k, v := range values {
		vars = append(vars, fmt.Sprintf("-var=%s=%s", k, v))
	}

	cmd := exec.Command("terraform", vars...)
	cmd.Dir = folder

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Dir = folder

	return cmd.Run()
}
