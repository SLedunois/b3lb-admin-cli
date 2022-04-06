package root

import (
	"fmt"
	"os"

	"github.com/SLedunois/b3lb/v2/pkg/restclient"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/SLedunois/b3lbctl/pkg/cmd/clusterinfo"
	configcmd "github.com/SLedunois/b3lbctl/pkg/cmd/config"
	"github.com/SLedunois/b3lbctl/pkg/cmd/instances"
	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/SLedunois/b3lbctl/pkg/system"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd struct represents the root command object
type RootCmd struct {
	Command *cobra.Command
	Flags   *Flags
}

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

	cmd := &RootCmd{
		Command: &cobra.Command{
			Use:   "b3lbctl <command> [flags]",
			Short: "B3LB controller cli",
			Long:  `Manage your B3LB cluster from the command line`,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					cmd.Help()
				}
			},
		},
		Flags: NewRootFlags(),
	}

	cmd.ApplyFlags()

	cmd.Command.AddCommand(instances.NewCmd())
	cmd.Command.AddCommand(clusterinfo.NewCmd())
	cmd.Command.AddCommand(configcmd.NewCmd())

	return cmd.Command
}
