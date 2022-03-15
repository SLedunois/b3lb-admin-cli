package config

// Flags contains all config command flags
type Flags struct{}

// NewFlags initialize a new ConfigFlags object
func NewFlags() *Flags {
	return &Flags{}
}

// ViewFlags contains all `config view` command flags
type ViewFlags struct{}

// NewViewFlags initialize a new ViewFlags object
func NewViewFlags() *ViewFlags {
	return &ViewFlags{}
}
