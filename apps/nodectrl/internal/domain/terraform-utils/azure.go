package terraformutils

type azureProvider struct {
	provider string
}

func NewAzureProvider() *azureProvider {
	return &azureProvider{
		provider: "",
	}
}
