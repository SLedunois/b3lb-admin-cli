package root

import (
	"fmt"
	"os"
	"strings"

	"github.com/SLedunois/b3lb/v2/pkg/restclient"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/SLedunois/b3lbctl/pkg/cmd/clusterinfo"
	configcmd "github.com/SLedunois/b3lbctl/pkg/cmd/describe"
	"github.com/SLedunois/b3lbctl/pkg/cmd/get"
	initcmd "github.com/SLedunois/b3lbctl/pkg/cmd/init"
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

// IsInitCommand returns true if the command is `b3lbctl init config`
func IsInitCommand() bool {
	return strings.Contains(strings.Join(os.Args, ""), "init")
}

// NewCmd initialize the root command
func NewCmd() *cobra.Command {
	cobra.OnInitialize(func() {
		restclient.Init()
		admin.Init()
		if !IsInitCommand() {
			err := config.Init(viper.GetString("config"))
			if err != nil {
				fmt.Println(err)
				os.Exit(system.NoSuchFileOrDirectoryExitCode)
			}
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

	cmd.Command.AddCommand(get.NewCmd())
	cmd.Command.AddCommand(clusterinfo.NewCmd())
	cmd.Command.AddCommand(configcmd.NewCmd())
	cmd.Command.AddCommand(initcmd.NewCmd())

	return cmd.Command
}
