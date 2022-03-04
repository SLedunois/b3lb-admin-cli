package instances

import (
	"github.com/spf13/cobra"
)

// NewCmd returns instances subcommand
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instances <command> [flags]",
		Short: "Manage B3LB instances",
		Long:  `Manage your B3LB cluster instances from the command line`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewAddCmd())

	return cmd
}
