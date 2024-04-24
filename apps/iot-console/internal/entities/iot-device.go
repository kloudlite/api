package entities

import (
	"fmt"
	"strings"

	// "github.com/kloudlite/api/apps/iot-console/internal/domain"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/common/fields"
	"github.com/kloudlite/api/constants"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
	rApi "github.com/kloudlite/operator/pkg/operator"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IOTDevice struct {
	repos.BaseEntity        `json:",inline" graphql:"noinput"`
	common.ResourceMetadata `json:",inline"`

	Name           string `json:"name"`
	AccountName    string `json:"accountName" graphql:"noinput"`
	ProjectName    string `json:"projectName" graphql:"noinput"`
	PublicKey      string `json:"publicKey"`
	ServiceCIDR    string `json:"serviceCIDR" graphql:"noinput"`
	PodCIDR        string `json:"podCIDR" graphql:"noinput"`
	IP             string `json:"ip" graphql:"noinput"`
	DeploymentName string `json:"deploymentName" graphql:"noinput"`

	ClusterToken string `json:"clusterToken" graphql:"noinput"`
	Index        int    `json:"index" graphql:"noinput"`

	CurrentBlueprint  *IOTDeviceBlueprint `json:"currentBlueprint,omitempty" graphql:"noinput"`
	ExpectedBlueprint *IOTDeviceBlueprint `json:"expectedBlueprint,omitempty" graphql:"noinput"`

	Status rApi.Status `json:"status,omitempty" graphql:"noinput"`

	SyncStatus t.SyncStatus `json:"syncStatus" graphql:"noinput"`
}

func (i *IOTDevice) GetName() string {
	return i.Name
}

func (i *IOTDevice) GetNamespace() string {
	return i.ProjectName
}

func (IOTDevice) GetCreationTimestamp() metav1.Time {
	return metav1.Now()
}

func (i *IOTDevice) GetLabels() map[string]string {
	return map[string]string{}
}

func (i *IOTDevice) GetDisplayName() string {
	return i.Name
}

func (i *IOTDevice) GetAnnotations() map[string]string {
	return map[string]string{}
}

func (i *IOTDevice) GetGeneration() int64 {
	return 1
}

func (i *IOTDevice) GetStatus() rApi.Status {
	return i.Status
}

func (i *IOTDevice) GetRecordVersion() int {
	return 1
}

// func (i *IOTDevice) GetResourceType() domain.ResourceType {
// 	return domain.ResourceTypeIOTDevice
// }

type DeviceWithServices struct {
	*IOTDevice
	ExposedDomains []string `json:"exposedDomains"`
	ExposedIps     []string `json:"exposedIps"`
}

func (i *IOTDevice) GetClusterName() string {
	return fmt.Sprintf("%s-%s", constants.DeviceClusterPrefix, i.Name)
}

func GetClusterName(deviceName string) string {
	return fmt.Sprintf("%s-%s", constants.DeviceClusterPrefix, deviceName)
}

func ExtractDeviceName(clusterName string) (*string, error) {
	var deviceName string
	prefix := fmt.Sprintf("%s-", constants.DeviceClusterPrefix)

	if strings.HasPrefix(clusterName, prefix) {
		deviceName = strings.TrimPrefix(clusterName, prefix)
	} else {
		return nil, errors.NewE(fmt.Errorf("cluster name %s does not start with %s", clusterName, prefix))
	}

	return &deviceName, nil
}

var IOTDeviceIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: fields.Id, Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "name", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "deploymentName", Value: repos.IndexAsc},
			{Key: "index", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
