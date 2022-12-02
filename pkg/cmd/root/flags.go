package root

import (
	"fmt"

	"github.com/bigblueswarm/bbsctl/pkg/config"
	"github.com/spf13/viper"
)

// Flags contains all root command flags
type Flags struct {
	ConfigPath string
}

// NewRootFlags initialize root command flags. It returns a RootFlags struct
func NewRootFlags() *Flags {
	return &Flags{}
}

// ApplyFlags apply RootFlags to provided command
func (cmd *RootCmd) ApplyFlags() {
	if !IsInitCommand() {
		cmd.Command.PersistentFlags().StringVar(&cmd.Flags.ConfigPath, "config", config.DefaultConfigPath(), fmt.Sprintf("config file (default is %s)", config.DefaultConfigPath()))
		cmd.Command.MarkFlagRequired("config")
		viper.BindPFlag("config", cmd.Command.PersistentFlags().Lookup("config"))
	}
}
