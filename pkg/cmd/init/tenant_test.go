package init

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/bigblueswarm/bbsctl/internal/test"
	bbsadmin "github.com/bigblueswarm/bigblueswarm/v2/pkg/admin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestInitTenantCmd(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
		return
	}

	os.Remove(fmt.Sprintf("%s/.bigblueswarm", homedir))

	tests := []test.CmdTest{
		{
			Name: "a valid command should init a new tenant file",
			Args: []string{"--host", "localhost"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				file := fmt.Sprintf("%s/.bigblueswarm/localhost.tenant.yml", homedir)
				defer os.Remove(file)

				b, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
					return
				}

				var tenant bbsadmin.Tenant
				if err := yaml.Unmarshal(b, &tenant); err != nil {
					t.Fatal(err)
					return
				}

				assert.Equal(t, "Tenant", tenant.Kind)
				assert.Equal(t, "localhost", tenant.Spec.Host)
				assert.Nil(t, tenant.Spec.MeetingsPool)
				assert.Nil(t, tenant.Spec.UserPool)
				assert.Equal(t, 0, len(tenant.Instances))
				assert.Nil(t, err)
				assert.Equal(t, fmt.Sprintf("tenant file successfully initialized. Please check %s/.bigblueswarm/localhost.tenant.yml file\n", homedir), string(output.Bytes()))
			},
		},
		{
			Name: "adding a meeting_pool flag should configure the tenant spec meeting pool configuration",
			Args: []string{"--host", "localhost", "--meeting_pool", "10"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				file := fmt.Sprintf("%s/.bigblueswarm/localhost.tenant.yml", homedir)
				defer os.Remove(file)

				b, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
					return
				}

				var tenant bbsadmin.Tenant
				if err := yaml.Unmarshal(b, &tenant); err != nil {
					t.Fatal(err)
					return
				}

				assert.Equal(t, "Tenant", tenant.Kind)
				assert.Equal(t, "localhost", tenant.Spec.Host)
				assert.Equal(t, int64(10), *tenant.Spec.MeetingsPool)
				assert.Nil(t, tenant.Spec.UserPool)
				assert.Equal(t, 0, len(tenant.Instances))
				assert.Nil(t, err)
				assert.Equal(t, fmt.Sprintf("tenant file successfully initialized. Please check %s/.bigblueswarm/localhost.tenant.yml file\n", homedir), string(output.Bytes()))
			},
		},
		{
			Name: "adding a user_pool flag should configure the tenant spec user pool configuration",
			Args: []string{"--host", "localhost", "--user_pool", "100"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				file := fmt.Sprintf("%s/.bigblueswarm/localhost.tenant.yml", homedir)
				defer os.Remove(file)

				b, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
					return
				}

				var tenant bbsadmin.Tenant
				if err := yaml.Unmarshal(b, &tenant); err != nil {
					t.Fatal(err)
					return
				}

				assert.Equal(t, "Tenant", tenant.Kind)
				assert.Equal(t, "localhost", tenant.Spec.Host)
				assert.Equal(t, int64(100), *tenant.Spec.UserPool)
				assert.Nil(t, tenant.Spec.MeetingsPool)
				assert.Equal(t, 0, len(tenant.Instances))
				assert.Nil(t, err)
				assert.Equal(t, fmt.Sprintf("tenant file successfully initialized. Please check %s/.bigblueswarm/localhost.tenant.yml file\n", homedir), string(output.Bytes()))
			},
		},
		{
			Name: "a valid command should init a new tenant file",
			Args: []string{"--host", "localhost"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				file := fmt.Sprintf("%s/.bigblueswarm/localhost.tenant.yml", homedir)
				b, err := os.ReadFile(file)
				if err != nil {
					t.Fatal(err)
					return
				}

				var tenant bbsadmin.Tenant
				if err := yaml.Unmarshal(b, &tenant); err != nil {
					t.Fatal(err)
					return
				}

				assert.Equal(t, "Tenant", tenant.Kind)
				assert.Equal(t, "localhost", tenant.Spec.Host)
				assert.Equal(t, 0, len(tenant.Instances))
				assert.Nil(t, err)
				assert.Equal(t, fmt.Sprintf("tenant file successfully initialized. Please check %s/.bigblueswarm/localhost.tenant.yml file\n", homedir), string(output.Bytes()))
			},
		},
		{
			Name: "initializing an existing tenant should return an error",
			Args: []string{"--host", "localhost"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("localhost.tenant.yml tenant file already exists. Please consider editing %s file", path.Join(homedir, ".bigblueswarm", "localhost.tenant.yml")), err.Error())
			},
		},
		{
			Name: "initializing a tenant in a folder that you do not have permsson should return an error",
			Args: []string{"--host", "localhost", "--dest", "/etc"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "failed to write tenant file: open /etc/localhost.tenant.yml: permission denied", err.Error())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewInitTenantCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
