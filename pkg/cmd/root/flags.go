package root

import (
	"fmt"

	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/spf13/viper"
)

// RootFlags contains all root command flags
type RootFlags struct {
	ConfigPath string
}

// NewRootFlags initialize root command flags. It returns a RootFlags struct
func NewRootFlags() *RootFlags {
	return &RootFlags{}
}

// ApplyFlags apply RootFlags to provided command
func (cmd *RootCmd) ApplyFlags() {
	cmd.Command.PersistentFlags().StringVar(&cmd.Flags.ConfigPath, "config", config.DefaultConfigPath(), fmt.Sprintf("config file (default is %s)", config.DefaultConfigPath()))
	cmd.Command.MarkFlagRequired("config")
	viper.BindPFlag("config", cmd.Command.PersistentFlags().Lookup("config"))
}
