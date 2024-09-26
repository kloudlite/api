package ports

import "context"

type InfraService interface {
	EnsureGlobalVPNConnection(ctx context.Context, args EnsureGlobalVPNConnectionIn) error
	GetClusterLabels(ctx context.Context, args IsClusterLabelsIn) (map[string]string, error)
}

type IsClusterLabelsIn struct {
	UserId    string
	UserEmail string
	UserName  string

	AccountName string
	ClusterName string
}

type EnsureGlobalVPNConnectionIn struct {
	UserId    string
	UserEmail string
	UserName  string

	AccountName   string
	ClusterName   string
	GlobalVPNName string

	DispatchAddrAccountName string
	DispatchAddrClusterName string
}
