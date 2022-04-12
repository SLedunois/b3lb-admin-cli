package describe

import (
	"bytes"
	"errors"
	"testing"

	b3lbadmin "github.com/SLedunois/b3lb/v2/pkg/admin"
	"gopkg.in/yaml.v3"

	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/SLedunois/b3lbctl/internal/test"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantCmd(t *testing.T) {
	assert.NotNil(t, NewTenantCmd())
}

func TestDescribeTenantC(t *testing.T) {
	admin.API = &mock.AdminMock{}

	tenant := &b3lbadmin.Tenant{
		Kind: "Tenant",
		Spec: map[string]string{
			"host": "localhost",
		},
		Instances: []string{},
	}

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
			Name: "an error returned by admin should return an error",
			Args: []string{"localhost"},
			Mock: func() {
				mock.GetTenantFunc = func(hostname string) (*b3lbadmin.Tenant, error) {
					return nil, errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to describe tenant localhost: admin error", err.Error())
			},
		},
		{
			Name: "a valid command should describe a valid tenant",
			Args: []string{"localhost"},
			Mock: func() {
				mock.GetTenantFunc = func(hostname string) (*b3lbadmin.Tenant, error) {
					return tenant, nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				var value *b3lbadmin.Tenant
				if err := yaml.Unmarshal(output.Bytes(), &value); err != nil {
					t.Fatal(err)
					return
				}

				assert.Nil(t, err)
				assert.Equal(t, tenant, value)
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
