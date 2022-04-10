package init

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	b3lbadmin "github.com/SLedunois/b3lb/v2/pkg/admin"
	"github.com/SLedunois/b3lbctl/internal/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestInitInstances(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
		return
	}

	os.Remove(fmt.Sprintf("%s/.b3lb", homedir))

	tests := []test.CmdTest{
		{
			Name: "a valid command should init instances configuration file",
			Mock: func() {},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				file := fmt.Sprintf("%s/.b3lb/instances.yml", homedir)
				b, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
					return
				}

				var instances b3lbadmin.InstanceList
				if err := yaml.Unmarshal(b, &instances); err != nil {
					t.Fatal(err)
					return
				}

				assert.Equal(t, "InstanceList", instances.Kind)
				assert.Equal(t, 0, len(instances.Instances))
			},
		},
		{
			Name: "an existing file should return an error",
			Mock: func() {},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "instances configuration file already exists. Please consider editing /home/codespace/.b3lb/instances.yml file", err.Error())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewInitInstancesCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			test.Mock()
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
