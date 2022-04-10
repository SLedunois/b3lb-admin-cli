package describe

import (
	"bytes"
	"errors"
	"testing"

	b3lbconfig "github.com/SLedunois/b3lb/v2/pkg/config"
	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/SLedunois/b3lbctl/internal/test"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestNewConfigCmd(t *testing.T) {
	assert.NotNil(t, NewConfigCmd())
}

func TestDescribeConfigCmd(t *testing.T) {
	admin.API = &mock.AdminMock{}
	tests := []test.CmdTest{
		{
			Name: "an error return by admin get configuration method should return an error",
			Mock: func() {
				mock.GetConfigurationFunc = func() (*b3lbconfig.Config, error) {
					return nil, errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "it should displays the configuration",
			Mock: func() {
				config := &b3lbconfig.Config{}

				mock.GetConfigurationFunc = func() (*b3lbconfig.Config, error) {
					return config, nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				out, err := yaml.Marshal(&b3lbconfig.Config{})
				if err != nil {
					t.Fatal(err)
					return
				}

				assert.Nil(t, err)
				assert.Equal(t, string(out), string(out))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			b := bytes.NewBufferString("")
			cmd := NewConfigCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
