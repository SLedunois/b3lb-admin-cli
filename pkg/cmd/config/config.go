package config

import "github.com/spf13/cobra"

// ConfigCmd represents config command type
type ConfigCmd struct {
	Flags   *ConfigFlags
	Command *cobra.Command
}

// NewCmd initialize a new config command
func NewCmd() *cobra.Command {
	cmd := &ConfigCmd{
		Flags: NewConfigFlags(),
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
