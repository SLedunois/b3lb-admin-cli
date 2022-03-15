package config

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/SLedunois/b3lbctl/internal/test"
	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNewViewCmd(t *testing.T) {
	assert.NotNil(t, NewViewCmd())
}

func TestProcess(t *testing.T) {
	tests := []test.CmdTest{
		{
			Name: "a non existing file should return an error",
			Mock: func() {
				path := "/tmp/config.yml"
				config.Path = &path
			},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an existing file should display the configuration",
			Mock: func() {
				path := "../../../test/config.yml"
				config.Path = &path
			},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := `bigbluebutton:
  secret: bbb_secret
  recordings_poll_interval: 1m
balancer:
  metrics_range: -5m
admin:
  api_key: b3lb_admin_api_key
  url: http://localhost:8090
port: 8090
redis:
  address: localhost:6379
  password:
  database: 0
influxdb:
  address: http://localhost:8086
  token: influxdb_token
  organization: b3lb
  bucket: bucket
`
				assert.Equal(t, expected, string(out))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			b := bytes.NewBufferString("")
			cmd := NewViewCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
