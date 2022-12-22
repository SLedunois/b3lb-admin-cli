// Package delete provides the delete command
package delete

import (
	"errors"
	"fmt"

	"github.com/bigblueswarm/bbsctl/pkg/admin"
	"github.com/spf13/cobra"
)

// TenantCmd represents delete tenant command type
type TenantCmd struct {
	Flags   *TenantFlags
	Command *cobra.Command
}

// NewTenantCmd initialize a new delete tenant cmmand
func NewTenantCmd() *cobra.Command {
	cmd := &TenantCmd{
		Command: &cobra.Command{
			Use:   "tenant <hostname>",
			Short: "delete a tenant based on hostname",
			Long:  "delete a tenant based on hostname",
		},
		Flags: NewTenantFlags(),
	}

	cmd.Command.PreRunE = cmd.prerun
	cmd.Command.RunE = cmd.process

	return cmd.Command
}

func (cmd *TenantCmd) prerun(command *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("command failed: hostname not found in arguments")
	}

	cmd.Flags.Hostname = args[0]
	command.SilenceUsage = true
	return nil
}

func (cmd *TenantCmd) process(command *cobra.Command, args []string) error {
	err := admin.API.DeleteTenant(cmd.Flags.Hostname)
	if err != nil {
		return err
	}

	command.Println(fmt.Sprintf("Tenant %s successfully deleted", cmd.Flags.Hostname))
	return nil
}
