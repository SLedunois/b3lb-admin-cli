package config

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
			Use:   "config <command> [flags]",
			Short: "Manage B3LB config file using subcommands.",
			Long:  `Manage B3LB config file using subcommands.`,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					cmd.Help()
				}
			},
		},
	}

	cmd.Command.AddCommand(NewViewCmd())

	return cmd.Command
}
