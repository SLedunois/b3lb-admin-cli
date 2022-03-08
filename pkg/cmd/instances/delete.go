package instances

import (
	"fmt"

	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/spf13/cobra"
)

// DeleteCmd struct represents Delete command object
type DeleteCmd struct {
	Command *cobra.Command
	Flags   *DeleteFlags
}

// NewDeleteCmd return the instance of delete subcommand
func NewDeleteCmd() *cobra.Command {
	cmd := &DeleteCmd{
		Command: &cobra.Command{
			Use:   "delete",
			Short: "Delete a BigBlueButton instance",
			Long:  `Delete a BigBlueButton instance from your B3LB cluster`,
		},
		Flags: NewDeleteFlags(),
	}

	cmd.Command.RunE = cmd.process
	cmd.ApplyFlags()
	return cmd.Command
}

func (cmd *DeleteCmd) process(command *cobra.Command, args []string) error {
	if err := admin.API.Delete(cmd.Flags.URL); err != nil {
		command.SilenceUsage = true
		return fmt.Errorf("unable to delete %s BigBlueButton instance from your cluster: %s", cmd.Flags.URL, err.Error())
	}

	command.Println(fmt.Sprintf(`instance %s deleted`, cmd.Flags.URL))
	return nil
}
