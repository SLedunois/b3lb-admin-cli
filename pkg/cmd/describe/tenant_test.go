package describe

import (
	"bytes"
	"errors"
	"testing"

	bbsadmin "github.com/bigblueswarm/bigblueswarm/v2/pkg/admin"
	"gopkg.in/yaml.v3"

	"github.com/bigblueswarm/bbsctl/internal/mock"
	"github.com/bigblueswarm/bbsctl/internal/test"
	"github.com/bigblueswarm/bbsctl/pkg/admin"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantCmd(t *testing.T) {
	assert.NotNil(t, NewTenantCmd())
}

func TestDescribeTenantC(t *testing.T) {
	admin.API = &mock.AdminMock{}

	tenant := &bbsadmin.Tenant{
		Kind: "Tenant",
		Spec: &bbsadmin.TenantSpec{
			Host: "localhost",
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
				mock.GetTenantFunc = func(hostname string) (*bbsadmin.Tenant, error) {
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
				mock.GetTenantFunc = func(hostname string) (*bbsadmin.Tenant, error) {
					return tenant, nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				var value *bbsadmin.Tenant
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
