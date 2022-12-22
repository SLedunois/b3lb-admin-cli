// Package delete provides the delete command
package delete

// Flags contains all delete command flags
type Flags struct{}

// NewFlags initialize a new Flags struct
func NewFlags() *Flags {
	return &Flags{}
}

// TenantFlags contains all delete tenant command flags
type TenantFlags struct {
	Hostname string
}

// NewTenantFlags initialize a new TenantFlags struct
func NewTenantFlags() *TenantFlags {
	return &TenantFlags{}
}
