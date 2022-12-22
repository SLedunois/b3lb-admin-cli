// Package get provides the get command
package get

// Flags contains all Get command flags
type Flags struct {
	CSV  bool
	JSON bool
}

// NewFlags initialize a new Flags object
func NewFlags() *Flags {
	return &Flags{
		CSV:  false,
		JSON: false,
	}
}
