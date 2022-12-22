// Package describe provides the describe command
package describe

// Flags contains all config command flags
type Flags struct{}

// NewFlags initialize a new ConfigFlags object
func NewFlags() *Flags {
	return &Flags{}
}

// ConfigFlags contains all `bbsctl describe config` command flags
type ConfigFlags struct{}

// NewConfigFlags initialize a new ConfigFlags object
func NewConfigFlags() *ConfigFlags {
	return &ConfigFlags{}
}

// TenantFlags contains all `bbsctl describe tenant <tenant> command flags`
type TenantFlags struct {
	Hostname string
}

// NewTenantFlags initialize a new TenantFlags object
func NewTenantFlags() *TenantFlags {
	return &TenantFlags{}
}
