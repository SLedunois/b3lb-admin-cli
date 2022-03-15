package config

// ConfigFlags contains all config command flags
type ConfigFlags struct{}

// NewConfigFlags initialize a new ConfigFlags object
func NewConfigFlags() *ConfigFlags {
	return &ConfigFlags{}
}

// ViewFlags contains all `config view` command flags
type ViewFlags struct{}

// NewViewFlags initialize a new ViewFlags object
func NewViewFlags() *ViewFlags {
	return &ViewFlags{}
}
