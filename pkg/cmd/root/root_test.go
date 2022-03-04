package root

import (
	"testing"

	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	type test struct {
		name      string
		args      []string
		validator func(t *testing.T, err error)
	}

	tests := []test{
		{
			name: "Root command should initialize configuration",
			args: []string{"--config", "../../../test/config.yml"},
			validator: func(t *testing.T, err error) {
				assert.Equal(t, "b3lb_admin_api_key", config.APIKey)
				assert.Equal(t, "http://localhost:8090", config.URL)
			},
		},
	}

	for _, test := range tests {
		cmd := NewCmd()
		cmd.SetArgs(test.args)
		cmd.Execute()
	}
}
