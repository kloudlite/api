// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type BYOCClusterSpec struct {
	AccountName         string    `json:"accountName"`
	DefaultIngressClass *string   `json:"defaultIngressClass,omitempty"`
	DefaultStorageClass *string   `json:"defaultStorageClass,omitempty"`
	Provider            string    `json:"provider"`
	PublicIps           []*string `json:"publicIps,omitempty"`
	Region              string    `json:"region"`
}

type BYOCClusterSpecIn struct {
	AccountName         string    `json:"accountName"`
	DefaultIngressClass *string   `json:"defaultIngressClass,omitempty"`
	DefaultStorageClass *string   `json:"defaultStorageClass,omitempty"`
	Provider            string    `json:"provider"`
	PublicIps           []*string `json:"publicIps,omitempty"`
	Region              string    `json:"region"`
}

type CloudProviderSpec struct {
	DisplayName    string                           `json:"display_name"`
	Provider       string                           `json:"provider"`
	ProviderSecret *CloudProviderSpecProviderSecret `json:"providerSecret"`
	AccountName    string                           `json:"accountName"`
}

type CloudProviderSpecIn struct {
	DisplayName    string                             `json:"display_name"`
	Provider       string                             `json:"provider"`
	ProviderSecret *CloudProviderSpecProviderSecretIn `json:"providerSecret"`
	AccountName    string                             `json:"accountName"`
}

type CloudProviderSpecProviderSecret struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type CloudProviderSpecProviderSecretIn struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type ClusterSpec struct {
	AccountName  string `json:"accountName"`
	Config       string `json:"config"`
	Count        int    `json:"count"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type ClusterSpecIn struct {
	AccountName  string `json:"accountName"`
	Config       string `json:"config"`
	Count        int    `json:"count"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type EdgeSpec struct {
	ClusterName  string           `json:"clusterName"`
	Pools        []*EdgeSpecPools `json:"pools,omitempty"`
	Provider     *string          `json:"provider,omitempty"`
	ProviderName string           `json:"providerName"`
	Region       string           `json:"region"`
	AccountName  string           `json:"accountName"`
}

type EdgeSpecIn struct {
	ClusterName  string             `json:"clusterName"`
	Pools        []*EdgeSpecPoolsIn `json:"pools,omitempty"`
	Provider     *string            `json:"provider,omitempty"`
	ProviderName string             `json:"providerName"`
	Region       string             `json:"region"`
	AccountName  string             `json:"accountName"`
}

type EdgeSpecPools struct {
	Config string `json:"config"`
	Max    *int   `json:"max,omitempty"`
	Min    *int   `json:"min,omitempty"`
	Name   string `json:"name"`
}

type EdgeSpecPoolsIn struct {
	Config string `json:"config"`
	Max    *int   `json:"max,omitempty"`
	Min    *int   `json:"min,omitempty"`
	Name   string `json:"name"`
}

type MasterNodeSpec struct {
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
}

type MasterNodeSpecIn struct {
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
}

type NodePoolSpec struct {
	Min          *int   `json:"min,omitempty"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	EdgeName     string `json:"edgeName"`
	Max          *int   `json:"max,omitempty"`
}

type NodePoolSpecIn struct {
	Min          *int   `json:"min,omitempty"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
	AccountName  string `json:"accountName"`
	ClusterName  string `json:"clusterName"`
	Config       string `json:"config"`
	EdgeName     string `json:"edgeName"`
	Max          *int   `json:"max,omitempty"`
}

type WorkerNodeSpec struct {
	EdgeName     string `json:"edgeName"`
	NodeIndex    *int   `json:"nodeIndex,omitempty"`
	Pool         string `json:"pool"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Stateful     *bool  `json:"stateful,omitempty"`
	Config       string `json:"config"`
	ClusterName  string `json:"clusterName"`
	Region       string `json:"region"`
	AccountName  string `json:"accountName"`
}

type WorkerNodeSpecIn struct {
	EdgeName     string `json:"edgeName"`
	NodeIndex    *int   `json:"nodeIndex,omitempty"`
	Pool         string `json:"pool"`
	Provider     string `json:"provider"`
	ProviderName string `json:"providerName"`
	Stateful     *bool  `json:"stateful,omitempty"`
	Config       string `json:"config"`
	ClusterName  string `json:"clusterName"`
	Region       string `json:"region"`
	AccountName  string `json:"accountName"`
}

type SyncState string

const (
	SyncStateIDLe       SyncState = "IDLE"
	SyncStateInProgress SyncState = "IN_PROGRESS"
	SyncStateReady      SyncState = "READY"
	SyncStateNotReady   SyncState = "NOT_READY"
)

var AllSyncState = []SyncState{
	SyncStateIDLe,
	SyncStateInProgress,
	SyncStateReady,
	SyncStateNotReady,
}

func (e SyncState) IsValid() bool {
	switch e {
	case SyncStateIDLe, SyncStateInProgress, SyncStateReady, SyncStateNotReady:
		return true
	}
	return false
}

func (e SyncState) String() string {
	return string(e)
}

func (e *SyncState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SyncState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SyncState", str)
	}
	return nil
}

func (e SyncState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
