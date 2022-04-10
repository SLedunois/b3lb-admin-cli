package init

// ConfigFlags represents `b3lbctl init config` flags
type ConfigFlags struct {
	B3LB        string
	APIKey      string
	Destination string
}

// NewInitConfigFlags initialize a new InitConfigFlags struct
func NewInitConfigFlags() *ConfigFlags {
	return &ConfigFlags{}
}

// InstancesFlags represents `b3lbctl init instances` flags
type InstancesFlags struct {
	Destination string
}

// NewInitInstancesFlags initialize a new InitInstancesFlags struct
func NewInitInstancesFlags() *InstancesFlags {
	return &InstancesFlags{}
}
