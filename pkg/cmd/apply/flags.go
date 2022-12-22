// Package apply provides apply command
package apply

// Flags contains all apply flags
type Flags struct {
	FilePath string
}

// NewFlags initialize a new Flags struct
func NewFlags() *Flags {
	return &Flags{}
}
