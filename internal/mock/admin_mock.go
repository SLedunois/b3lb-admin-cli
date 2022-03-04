package mock

import (
	"github.com/SLedunois/b3lb-admin-cli/pkg/admin"
	"github.com/SLedunois/b3lb/pkg/api"
)

var (
	// ListAdminFunc is the function that will be called when the mock admin is used
	ListAdminFunc func() ([]api.BigBlueButtonInstance, error)
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
