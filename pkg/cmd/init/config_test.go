package init

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/SLedunois/b3lbctl/internal/test"
	"github.com/SLedunois/b3lbctl/pkg/config"
	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
		return
	}

	os.Remove(fmt.Sprintf("%s/.b3lb", homedir))

	tests := []test.CmdTest{
		{
			Name: "a valid command should init configuration with parameters",
			Mock: func() {},
			Args: []string{"-b", "http://localhost:8090", "-k", "api_key"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				file := fmt.Sprintf("%s/.b3lb/.b3lbctl.yml", homedir)
				b, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
					return
				}

				var conf config.Config
				if err := yaml.Unmarshal(b, &conf); err != nil {
					t.Fatal(err)
					return
				}

				assert.Equal(t, "http://localhost:8090", conf.B3lb)
				assert.Equal(t, "api_key", conf.APIKey)
			},
		},
		{
			Name: "an existing file should return an `existing file` error",
			Mock: func() {},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("configuration already exists, see %s/.b3lb/.b3lbctl.yml", homedir), err.Error())
			},
		},
		{
			Name: "an error should be returned if I can't create a new folder",
			Mock: func() {},
			Args: []string{"-d", "/etc/b3lb"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to create destination folder: mkdir /etc/b3lb: permission denied", err.Error())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewInitConfigCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			test.Mock()
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
