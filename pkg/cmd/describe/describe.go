package describe

import "github.com/spf13/cobra"

// Cmd represents config command type
type Cmd struct {
	Flags   *Flags
	Command *cobra.Command
}

// NewCmd initialize a new config command
func NewCmd() *cobra.Command {
	cmd := &Cmd{
		Flags: NewFlags(),
		Command: &cobra.Command{
			Use:   "describe <command> [flags]",
			Short: "Show details of a specific resource or group of resources",
			Long:  `Show details of a specific resource or group of resources`,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					cmd.Help()
				}
			},
		},
	}

	cmd.Command.AddCommand(NewConfigCmd())

	return cmd.Command
}
