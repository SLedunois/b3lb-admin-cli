package describe

import (
	"errors"
	"fmt"

	"github.com/bigblueswarm/bbsctl/pkg/admin"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// TenantCmd represents describe tenant command
type TenantCmd struct {
	Command *cobra.Command
	Flags   *TenantFlags
}

// NewTenantCmd initialize a new TenantCmd
func NewTenantCmd() *cobra.Command {
	cmd := &TenantCmd{
		Command: &cobra.Command{
			Use:   "tenant <hostname>",
			Short: "Describe B3L tenant.",
			Long:  "Describe a given BigBlueSwarm tenant.",
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
	tenant, err := admin.API.GetTenant(cmd.Flags.Hostname)
	if err != nil {
		return fmt.Errorf("unable to describe tenant %s: %s", cmd.Flags.Hostname, err.Error())
	}

	out, err := yaml.Marshal(tenant)
	if err != nil {
		return fmt.Errorf("unable to describe tenant %s: %s", cmd.Flags.Hostname, err.Error())
	}

	command.Println(normalizeYaml(string(out)))

	return nil
}
