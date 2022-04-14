package delete

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/SLedunois/b3lbctl/internal/test"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantCmd(t *testing.T) {
	assert.NotNil(t, NewTenantCmd())
}

func TestDeleteTenant(t *testing.T) {
	admin.API = &mock.AdminMock{}
	tests := []test.CmdTest{
		{
			Name: "no tenant hostname in arguments should return an error",
			Args: []string{},
			Mock: func() {},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an error returned by admin should be returned",
			Args: []string{"localhost"},
			Mock: func() {
				mock.DeleteTenantFunc = func(hostname string) error {
					return errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "admin error", err.Error())
			},
		},
		{
			Name: "a valid command should return no error and display a valid message",
			Args: []string{"localhost"},
			Mock: func() {
				mock.DeleteTenantFunc = func(hostname string) error {
					return nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "Tenant localhost successfully deleted", strings.TrimSpace(string(output.Bytes())))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			b := bytes.NewBufferString("")
			cmd := NewTenantCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
