package get

// Flags contains all Get command flags
type Flags struct {
	CSV  bool
	JSON bool
}

// NewGetFlags initialize a new InstancesFlags object
func NewFlags() *Flags {
	return &Flags{
		CSV:  false,
		JSON: false,
	}
}
