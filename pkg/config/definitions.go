// Package config manages the configuration
package config

// Config represents the config file struct
type Config struct {
	BBS    string `yaml:"bbs"`
	APIKey string `yaml:"apiKey"`
}
