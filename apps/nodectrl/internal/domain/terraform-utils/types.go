package terraformutils

type ProviderClient interface {
	NewNode() error
	DeleteNode() error

	AttachNode() error
	UnattachNode() error
}
