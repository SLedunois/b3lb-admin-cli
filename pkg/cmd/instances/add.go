package instances

import (
	"fmt"

	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/spf13/cobra"
)

type AddCmd struct {
	Command *cobra.Command
	Flags   *AddFlags
}

// NewAddCmd return the instances add subcommand
func NewAddCmd() *cobra.Command {
	cmd := &AddCmd{
		Command: &cobra.Command{
			Use:   "add",
			Short: "Add a BigBlueButton instance",
			Long:  `Add a BigBlueButton instance in your B3LB cluster`,
		},
		Flags: NewAddFlags(),
	}

	cmd.Command.RunE = cmd.process

	cmd.ApplyFlags()

	return cmd.Command
}

func (cmd *AddCmd) process(command *cobra.Command, args []string) error {
	if err := admin.API.Add(cmd.Flags.URL, cmd.Flags.Secret); err != nil {
		command.SilenceUsage = true
		return fmt.Errorf("unable to add BigBlueButton instance to your cluster: %s", err.Error())
	}

	command.Println(fmt.Sprintf(`instance %s added`, cmd.Flags.URL))
	return nil
}
