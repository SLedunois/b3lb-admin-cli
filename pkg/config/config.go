package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SLedunois/b3lb/pkg/config"
)

const defaultConfigFileName = ".b3lb.yaml"

var defaultConfigPath = fmt.Sprintf("$HOME/%s", defaultConfigFileName)

// APIKey is the admin API key configuration found in configuration file
var APIKey *string

// URL is the admin url configuration found in configuration file
var URL *string

// Init initialize config and expose it through Config variable
func Init(path string) error {
	if path == defaultConfigPath {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		path = filepath.Join(homeDir, defaultConfigFileName)
	}

	c, err := config.Load(path)
	if err != nil {
		return fmt.Errorf("unable to load configuration: %s", err.Error())
	}

	APIKey = &c.Admin.APIKey
	URL = &c.Admin.URL

	return nil
}

// DefaultConfigPath returns the default config file path
func DefaultConfigPath() string {
	return defaultConfigPath
}
