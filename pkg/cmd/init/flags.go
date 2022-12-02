package init

// ConfigFlags represents `bbsctl init config` flags
type ConfigFlags struct {
	BBS         string
	APIKey      string
	Destination string
}

// NewInitConfigFlags initialize a new ConfigFlags struct
func NewInitConfigFlags() *ConfigFlags {
	return &ConfigFlags{}
}

// InstancesFlags represents `bbsctl init instances` flags
type InstancesFlags struct {
	Destination string
}

// NewInitInstancesFlags initialize a new InstancesFlags struct
func NewInitInstancesFlags() *InstancesFlags {
	return &InstancesFlags{}
}

// TenantFlags represents `bbsctl init tenant` flags
type TenantFlags struct {
	Hostname    string
	Destination string
}

// NewInitTenantFlags initialize a new TenantFlags struct
func NewInitTenantFlags() *TenantFlags {
	return &TenantFlags{}
}
