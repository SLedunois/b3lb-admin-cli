package init

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/bigblueswarm/bbsctl/internal/test"
	bbsadmin "github.com/bigblueswarm/bigblueswarm/v2/pkg/admin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestInitInstances(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
		return
	}

	os.Remove(fmt.Sprintf("%s/.bigblueswarm", homedir))

	tests := []test.CmdTest{
		{
			Name: "a valid command should init instances configuration file",
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				file := fmt.Sprintf("%s/.bigblueswarm/instances.yml", homedir)
				b, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
					return
				}

				var instances bbsadmin.InstanceList
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
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				home, _ := os.UserHomeDir()
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("instances configuration file already exists. Please consider editing %s/.bigblueswarm/instances.yml file", home), err.Error())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewInitInstancesCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
