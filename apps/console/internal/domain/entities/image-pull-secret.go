package entities

import (
	"encoding/json"
	"fmt"

	"encoding/base64"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/repos"
)

type ImagePullSecret struct {
	repos.BaseEntity `json:",inline" graphql:"noinput"`
	Name             string `json:"name"`
	AccountName      string `json:"accountName" grapqhl:"noinput"`

	DockerConfigJson *string `json:"dockerConfigJson,omitempty"`

	DockerUsername         *string `json:"dockerUsername,omitempty"`
	DockerPassword         *string `json:"dockerPassword,omitempty"`
	DockerRegistryEndpoint *string `json:"dockerRegistryEndpoint,omitempty"`
}

func (ips *ImagePullSecret) ToK8sSecret() (*corev1.Secret, error) {
	if ips.AccountName == "" {
		return nil, fmt.Errorf("account name is required")
	}

	dockerConfigJson, err := func() (*string, error) {
		if ips.DockerConfigJson != nil {
			return ips.DockerConfigJson, nil
		}

		b, err := json.Marshal(map[string]any{
			"auths": map[string]any{
				*ips.DockerRegistryEndpoint: map[string]any{
					"auth": base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", *ips.DockerUsername, *ips.DockerPassword))),
				},
			},
		})
		if err != nil {
			return nil, err
		}

		return fn.New(string(b)), nil
	}()

	if err != nil {
		return nil, err
	}

	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ips.Name,
			Namespace: ips.AccountName,
		},
		Immutable: fn.New(false),
		StringData: map[string]string{
			".dockerconfigjson": *dockerConfigJson,
		},
		Type: "kubernetes.io/dockerconfigjson",
	}, nil
}
