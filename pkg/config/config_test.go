package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	type test struct {
		name      string
		path      string
		validator func(t *testing.T, err error)
		precheck  func()
	}

	tests := []test{
		{
			name: "loading configuration should return an error if configuration is not found",
			path: DefaultConfigPath(),
			validator: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an unvailable $HOME environment variable should return an error",
			path: DefaultConfigPath(),
			validator: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
			precheck: func() {
				os.Unsetenv("HOME")
			},
		},
		{
			name: "a valid path should parse configuration",
			path: "../../test/config.yml",
			validator: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.precheck != nil {
				test.precheck()
			}

			test.validator(t, Init(test.path))
		})
	}
}
