package delete

import "github.com/spf13/cobra"

// Cmd represents delete command type
type Cmd struct {
	Command *cobra.Command
	Flags   *Flags
}

// NewCmd initialize a new delete command
func NewCmd() *cobra.Command {
	cmd := &Cmd{
		Command: &cobra.Command{
			Use:   "delete <command>",
			Short: "Delete a specific resource",
			Long:  "Delete a specific resource",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					cmd.Help()
				}
			},
		},
		Flags: NewFlags(),
	}

	cmd.Command.AddCommand(NewTenantCmd())

	return cmd.Command
}
