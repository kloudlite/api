package domain

type ProviderClient interface {
	NewNode() error
	DeleteNode() error

	AttachNode() error
	UnattachNode() error
}
