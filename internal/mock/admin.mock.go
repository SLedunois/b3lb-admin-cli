// Package mock provide an internal mock to test
package mock

import (
	"github.com/bigblueswarm/bbsctl/pkg/admin"
	bbsadmin "github.com/bigblueswarm/bigblueswarm/v2/pkg/admin"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/api"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/balancer"
	bbsconfig "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
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
	// BBSAPIStatusAdminFunc is the function that will be called when the mock admin is used
	BBSAPIStatusAdminFunc func() (string, error)
	// GetConfigurationFunc is the function that will be called when the mock admin is used
	GetConfigurationFunc func() (*bbsconfig.Config, error)
	// GetTenantsFunc is the function that will be called when the mock admin is used
	GetTenantsFunc func() (*bbsadmin.TenantList, error)
	// GetTenantFunc is the function that will be called when the mock admin is used
	GetTenantFunc func(hostname string) (*bbsadmin.Tenant, error)
	// DeleteTenantFunc is the function that will be called whtne the admin mock is used
	DeleteTenantFunc func(hostname string) error
	// ApplyFunc is the function that will be called when the admin mock is used
	ApplyFunc func(kind string, resource *interface{}) error
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

// Add is a mock implementation that add a bigbluebutton instance on BigBlueSwarm
func (a *AdminMock) Add(url string, secret string) error {
	return AddAdminFunc(url, secret)
}

// Delete is a mock implementation deleting a bigbluebutton instance on BigBlueSwarm
func (a *AdminMock) Delete(url string) error {
	return DeleteAdminFunc(url)
}

// ClusterStatus is a mock implementation returning a list of InstanceStatus
func (a *AdminMock) ClusterStatus() ([]balancer.InstanceStatus, error) {
	return ClusterStatusAdminFunc()
}

// BBSAPIStatus is a mock implementation returning a list of InstanceStatus
func (a *AdminMock) BBSAPIStatus() (string, error) {
	return BBSAPIStatusAdminFunc()
}

// GetConfiguration is a mock implementation returning the configuration
func (a *AdminMock) GetConfiguration() (*bbsconfig.Config, error) {
	return GetConfigurationFunc()
}

// GetTenants is a mock implementation return a TenantList
func (a *AdminMock) GetTenants() (*bbsadmin.TenantList, error) {
	return GetTenantsFunc()
}

// GetTenant is a mock implementation that return a given tenant as kind Tenant
func (a *AdminMock) GetTenant(hostname string) (*bbsadmin.Tenant, error) {
	return GetTenantFunc(hostname)
}

// DeleteTenant is a mock implement that delete a kind Tenant based on hostname
func (a *AdminMock) DeleteTenant(hostname string) error {
	return DeleteTenantFunc(hostname)
}

// Apply is a mock implementation that apply a resource
func (a *AdminMock) Apply(kind string, resource *interface{}) error {
	return ApplyFunc(kind, resource)
}
