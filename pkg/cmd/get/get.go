package get

import (
	"github.com/spf13/cobra"
)

// NewCmd returns instances subcommand
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <command> [flags]",
		Short: "Display a resource",
		Long:  `Display a resource`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			}
		},
	}

	cmd.AddCommand(NewInstancesCmd())
	cmd.AddCommand(NewTenantCmd())

	return cmd
}
