package mock

import (
	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lbctl/pkg/admin"
)

var (
	// ListAdminFunc is the function that will be called when the mock admin is used
	ListAdminFunc func() ([]api.BigBlueButtonInstance, error)
	// AddAdminFunc is the function that will be called when the mock admin is used
	AddAdminFunc func(url string, secret string) error
	//DeleteAdminFunc is the function that will be called when the mock admin is used
	DeleteAdminFunc func(url string) error
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
