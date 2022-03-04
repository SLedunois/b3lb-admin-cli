package root

import (
	"fmt"
	"os"

	"github.com/SLedunois/b3lb-admin-cli/pkg/admin"
	"github.com/SLedunois/b3lb-admin-cli/pkg/cmd/instances"
	"github.com/SLedunois/b3lb-admin-cli/pkg/config"
	"github.com/SLedunois/b3lb-admin-cli/pkg/system"
	"github.com/SLedunois/b3lb/pkg/restclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath string

// NewCmd initialize the root command
func NewCmd() *cobra.Command {
	cobra.OnInitialize(func() {
		restclient.Init()
		admin.Init()
		err := config.Init(viper.GetString("config"))
		if err != nil {
			fmt.Println(err)
			os.Exit(system.NoSuchFileOrDirectoryExitCode)
		}
	})

	cmd := &cobra.Command{
		Use:   "b3lb-admin <command> [flags]",
		Short: "B3LB admin cli",
		Long:  `Manage your B3LB cluster from the command line`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.PersistentFlags().StringVar(&configPath, "config", config.DefaultConfigPath(), fmt.Sprintf("config file (default is %s)", config.DefaultConfigPath()))
	viper.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))
	cmd.AddCommand(instances.NewCmd())

	return cmd
}
