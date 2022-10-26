package entities

import "kloudlite.io/pkg/repos"

type ManagedServiceType string
type ManagedResourceType string

type ManagedServiceCategory struct {
	Category    repos.ID                  `yaml:"category"json:"category"`
	LogoUrl     string                    `yaml:"logoUrl" json:"logo_url"`
	DisplayName string                    `yaml:"displayName"json:"display_name"`
	List        []*ManagedServiceTemplate `yaml:"list" json:"list"`
	Description string                    `yaml:"description"json:"description"`
}

type ManagedServiceTemplate struct {
	Name            string                    `yaml:"name" json:"name"`
	ApiVersion      string                    `yaml:"apiVersion" json:"api_version"`
	Kind            string                    `yaml:"kind" json:"kind"`
	LogoUrl         string                    `yaml:"logoUrl" json:"logo_url"`
	DisplayName     string                    `yaml:"displayName" json:"display_name"`
	Fields          []TemplateField           `yaml:"fields" json:"fields"`
	Outputs         []TemplateOutput          `yaml:"outputs" json:"outputs"`
	Resources       []ManagedResourceTemplate `yaml:"resources" json:"resources"`
	Active          bool                      `yaml:"active" json:"active"`
	Description     string                    `yaml:"description" json:"description"`
	InputMiddleware string                    `yaml:"inputMiddleware" json:"inputMiddleware"`
	Estimator       string                    `yaml:"estimator" json:"estimator"`
}

type TemplateOutput struct {
	Name  string `yaml:"name" json:"name"`
	Label string `yaml:"label" json:"label"`
}

type TemplateField struct {
	Name         string  `yaml:"name" json:"name"`
	Label        string  `yaml:"label" json:"label"`
	DisplayName  string  `yaml:"displayName" json:"display_name"`
	Description  string  `yaml:"description" json:"description"`
	Min          float32 `yaml:"min" json:"min"`
	Max          float32 `yaml:"max" json:"max"`
	DefaultValue string  `yaml:"defaultValue" json:"default_value"`
	Hidden       bool    `yaml:"hidden" json:"hidden"`
	InputType    string  `yaml:"inputType" json:"input_type"`
	Unit         string  `yaml:"unit" json:"unit"`
	Required     bool    `yaml:"required" json:"required"`
	Step         float32 `yaml:"step" json:"step"`
}

type ManagedResourceTemplate struct {
	Name        string           `yaml:"name" json:"name"`
	ApiVersion  string           `yaml:"apiVersion" json:"api_version"`
	Kind        string           `yaml:"kind" json:"kind"`
	DisplayName string           `yaml:"displayName" json:"display_name"`
	Fields      []*TemplateField `yaml:"fields" json:"fields"`
	Outputs     []TemplateOutput `yaml:"outputs" json:"outputs"`
	Default     bool             `json:"default,omitempty" yaml:"default,omitempty"`
	GetRefKey   string           `json:"getRefKey,omitempty" yaml:"getRefKey,omitempty"`
}
