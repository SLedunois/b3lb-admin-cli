// Package init provides the init command
package init

import (
	"github.com/spf13/cobra"
)

// NewCmd initialize a new init cobra command
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init <command> [flags]",
		Short: "Initialize a resource",
		Long:  "Initialize a resource",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			}
		},
	}

	cmd.AddCommand(NewInitConfigCmd())
	cmd.AddCommand(NewInitInstancesCmd())
	cmd.AddCommand(NewInitTenantCmd())

	return cmd
}
