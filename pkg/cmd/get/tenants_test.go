package get

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	b3lbadmin "github.com/SLedunois/b3lb/v2/pkg/admin"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"

	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/SLedunois/b3lbctl/internal/test"
)

func TestGetTenantsCmd(t *testing.T) {
	expected := &b3lbadmin.TenantList{
		Kind: "TenantList",
		Tenants: []b3lbadmin.TenantListObject{
			{
				Hostname:      "localhost",
				InstanceCount: 1,
			},
		},
	}

	mock.InitAdminMock()

	tests := []test.CmdTest{
		{
			Name: "an error thrown by admin should return an error",
			Mock: func() {
				mock.GetTenantsFunc = func() (*b3lbadmin.TenantList, error) {
					return nil, errors.New("admin error")
				}
			},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "calling get tenants cmd with --json sould print result as json response",
			Mock: func() {
				mock.GetTenantsFunc = func() (*b3lbadmin.TenantList, error) {
					return expected, nil
				}
			},
			Args: []string{"--json"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				assert.Equal(t, text.NewJSONTransformer("", "  ")(expected.Tenants), strings.TrimSpace(string(out)))
			},
		},
		{
			Name: "calling get tenants cmd with --csv sould print result as csv response",
			Mock: func() {
				mock.GetTenantsFunc = func() (*b3lbadmin.TenantList, error) {
					return expected, nil
				}
			},
			Args: []string{"--csv"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := `Hostname,Instances
localhost,1`
				assert.Equal(t, expected, strings.TrimSpace(string(out)))
			},
		},
		{
			Name: "calling get tenants cmd without flag sould print result as table response",
			Mock: func() {
				mock.GetTenantsFunc = func() (*b3lbadmin.TenantList, error) {
					return expected, nil
				}
			},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := `Hostname   Instances  
localhost          1`
				assert.Equal(t, expected, strings.TrimSpace(string(out)))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewTenantCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			test.Mock()
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
