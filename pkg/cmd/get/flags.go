package get

// InstancesFlags contains all List command flags
type InstancesFlags struct {
	CSV  bool
	JSON bool
}

// NewInstancesFlags initialize a new ListFlags object
func NewInstancesFlags() *InstancesFlags {
	return &InstancesFlags{
		CSV:  false,
		JSON: false,
	}
}
