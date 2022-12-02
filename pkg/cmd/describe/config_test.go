package describe

import (
	"bytes"
	"errors"
	"testing"

	"github.com/bigblueswarm/bbsctl/internal/mock"
	"github.com/bigblueswarm/bbsctl/internal/test"
	"github.com/bigblueswarm/bbsctl/pkg/admin"
	bbsconfig "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
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
				mock.GetConfigurationFunc = func() (*bbsconfig.Config, error) {
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
				config := &bbsconfig.Config{}

				mock.GetConfigurationFunc = func() (*bbsconfig.Config, error) {
					return config, nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				out, err := yaml.Marshal(&bbsconfig.Config{})
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
