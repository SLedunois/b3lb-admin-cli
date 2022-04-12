package mock

import (
	b3lbadmin "github.com/SLedunois/b3lb/v2/pkg/admin"
	"github.com/SLedunois/b3lb/v2/pkg/api"
	"github.com/SLedunois/b3lb/v2/pkg/balancer"
	b3lbconfig "github.com/SLedunois/b3lb/v2/pkg/config"
	"github.com/SLedunois/b3lbctl/pkg/admin"
)

var (
	// ListAdminFunc is the function that will be called when the mock admin is used
	ListAdminFunc func() ([]api.BigBlueButtonInstance, error)
	// AddAdminFunc is the function that will be called when the mock admin is used
	AddAdminFunc func(url string, secret string) error
	//DeleteAdminFunc is the function that will be called when the mock admin is used
	DeleteAdminFunc func(url string) error
	// ClusterStatusAdminFunc is the function that will be called when the mock admin is used
	ClusterStatusAdminFunc func() ([]balancer.InstanceStatus, error)
	// B3lbAPIStatusAdminFunc is the function that will be called when the mock admin is used
	B3lbAPIStatusAdminFunc func() (string, error)
	// GetConfigurationFunc is the function that will be called when the mock admin is used
	GetConfigurationFunc func() (*b3lbconfig.Config, error)
	// GetTenantsFunc is the function that will be called when the mock admin is used
	GetTenantsFunc func() (*b3lbadmin.TenantList, error)
	// GetTenantFunc is the function that will be called when the mock admin is used
	GetTenantFunc func(hostname string) (*b3lbadmin.Tenant, error)
)

// InitAdminMock init admin.API object with an empty AdminMock struct
func InitAdminMock() {
	admin.API = &AdminMock{}
}

// AdminMock represents an admin mock object
type AdminMock struct{}

// List is a mock implementation that return a list of all instances
func (a *AdminMock) List() ([]api.BigBlueButtonInstance, error) {
	return ListAdminFunc()
}

// Add is a mock implementation that add a bigbluebutton instance on b3lb
func (a *AdminMock) Add(url string, secret string) error {
	return AddAdminFunc(url, secret)
}

// Delete is a mock implementation deleting a bigbluebutton instance on b3lb
func (a *AdminMock) Delete(url string) error {
	return DeleteAdminFunc(url)
}

// ClusterStatus is a mock implementation returning a list of InstanceStatus
func (a *AdminMock) ClusterStatus() ([]balancer.InstanceStatus, error) {
	return ClusterStatusAdminFunc()
}

// B3lbAPIStatus is a mock implementation returning a list of InstanceStatus
func (a *AdminMock) B3lbAPIStatus() (string, error) {
	return B3lbAPIStatusAdminFunc()
}

// GetConfigurations is a mock implementation returning the configuration
func (a *AdminMock) GetConfiguration() (*b3lbconfig.Config, error) {
	return GetConfigurationFunc()
}

// GetTenants is a mock implementation return a TenantList
func (a *AdminMock) GetTenants() (*b3lbadmin.TenantList, error) {
	return GetTenantsFunc()
}

// GetTenant is a mock implementation that return a given tenant as kind Tenant
func (a *AdminMock) GetTenant(hostname string) (*b3lbadmin.Tenant, error) {
	return GetTenantFunc(hostname)
}
