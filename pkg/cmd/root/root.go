// Package root provide the root bbsctl command
package root

import (
	"fmt"
	"os"
	"strings"

	"github.com/bigblueswarm/bbsctl/pkg/admin"
	"github.com/bigblueswarm/bbsctl/pkg/cmd/apply"
	"github.com/bigblueswarm/bbsctl/pkg/cmd/clusterinfo"
	"github.com/bigblueswarm/bbsctl/pkg/cmd/delete"
	configcmd "github.com/bigblueswarm/bbsctl/pkg/cmd/describe"
	"github.com/bigblueswarm/bbsctl/pkg/cmd/get"
	initcmd "github.com/bigblueswarm/bbsctl/pkg/cmd/init"
	"github.com/bigblueswarm/bbsctl/pkg/config"
	"github.com/bigblueswarm/bbsctl/pkg/system"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/restclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd struct represents the root command object
type RootCmd struct {
	Command *cobra.Command
	Flags   *Flags
}

// IsInitCommand returns true if the command is `bbsctl init config`
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
			Use:   "bbsctl <command> [flags]",
			Short: "BigBlueSwarm controller cli",
			Long:  `Manage your BigBlueSwarm cluster from the command line`,
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
	cmd.Command.AddCommand(delete.NewCmd())
	cmd.Command.AddCommand(apply.NewCmd())

	return cmd.Command
}
