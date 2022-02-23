package root

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SLedunois/b3lb/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultConfigFileName = ".b3lb.yaml"

var defaultConfigPath = fmt.Sprintf("$HOME/%s", defaultConfigFileName)
var configPath string
var Config *config.AdminConfig

func NewCmdRoot() *cobra.Command {
	cobra.OnInitialize(initConfig)
	rootCmd := &cobra.Command{
		Use:   "b3lb-admin <command> [flags]",
		Short: "B3LB admin cli",
		Long:  `Manage your B3LB cluster from the command line`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.PersistentFlags().StringVar(&configPath, "config", "c", fmt.Sprintf("config file (default is %s)", defaultConfigPath))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.SetDefault("config", "$HOME/.b3lb.yaml")

	return rootCmd
}

func initConfig() {
	path := viper.GetString("config")
	if path == defaultConfigPath {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		path = filepath.Join(homeDir, defaultConfigFileName)
	}

	c, err := config.Load(path)
	if err != nil {
		fmt.Println(fmt.Errorf("unable to load configuration: %s", err.Error()))
		os.Exit(1)
	}

	Config = &c.Admin
}
