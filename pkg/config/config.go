// Package config manages the configuration
package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
	"gopkg.in/yaml.v2"
)

// DefaultConfigFileName is the default config file name
const DefaultConfigFileName = ".bbsctl.yml"

var defaultConfigPath = fmt.Sprintf("%s/%s", config.DefaultConfigFolder, DefaultConfigFileName)

// APIKey is the admin API key configuration found in configuration file
var APIKey *string

// BBS is the admin url configuration found in configuration file
var BBS *string

// Path is the direct path for configuration file
var Path *string

// Init initialize config and expose it through Config variable
func Init(path string) error {
	if path == defaultConfigPath {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		path = filepath.Join(homeDir, ".bigblueswarm", DefaultConfigFileName)
	}

	c, err := load(path)
	if err != nil {
		return fmt.Errorf("unable to load configuration: %s", err.Error())
	}

	BBS = &c.BBS
	APIKey = &c.APIKey
	Path = &path

	return nil
}

func load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(fmt.Sprintf("unable to close config file: %s", err))
		}
	}()

	conf := &Config{}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}

// DefaultConfigPath returns the default config file path
func DefaultConfigPath() string {
	return defaultConfigPath
}
