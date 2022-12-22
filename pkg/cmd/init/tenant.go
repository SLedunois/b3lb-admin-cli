// Package init provides the init command
package init

import (
	"fmt"
	"os"
	"path"

	bbsadmin "github.com/bigblueswarm/bigblueswarm/v2/pkg/admin"
	bbsconfig "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const tenantFileNameFormatter = "%s.tenant.yml"

// TenantCmd represents the `bbsctl init tenant` command
type TenantCmd struct {
	Command *cobra.Command
	Flags   *TenantFlags
}

// NewInitTenantCmd initialize a new TenantCmd
func NewInitTenantCmd() *cobra.Command {
	cmd := &TenantCmd{
		Command: &cobra.Command{
			Use:   "tenant [flags]",
			Short: "Initialize a new bigblueswarm tenant configuration file",
			Long:  "Initialize a new bigblueswarm tenant configuration file if not exits",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					cmd.Help()
				}
			},
		},
		Flags: NewInitTenantFlags(),
	}

	cmd.ApplyFlags()

	cmd.Command.PreRunE = cmd.preRunE
	cmd.Command.RunE = cmd.init

	return cmd.Command
}

// ApplyFlags apply command flags to InitInstancesCmd
func (cmd *TenantCmd) ApplyFlags() {
	cmd.Command.Flags().StringVarP(&cmd.Flags.Destination, "dest", "d", bbsconfig.DefaultConfigFolder, "File folder destination")
	cmd.Command.Flags().StringVarP(&cmd.Flags.Hostname, "host", "", "", "Tenant hostname")
	cmd.Command.MarkFlagRequired("host")
}

func (cmd *TenantCmd) init(command *cobra.Command, args []string) error {
	filename := fmt.Sprintf(tenantFileNameFormatter, cmd.Flags.Hostname)
	destFile := path.Join(cmd.Flags.Destination, filename)
	if fileAlreadyExists(destFile) {
		return fmt.Errorf("%s tenant file already exists. Please consider editing %s file", filename, destFile)
	}

	tenant := &bbsadmin.Tenant{
		Kind: "Tenant",
		Spec: &bbsadmin.TenantSpec{
			Host: cmd.Flags.Hostname,
		},
		Instances: []string{},
	}

	if err := os.MkdirAll(cmd.Flags.Destination, fsRights); err != nil {
		return fmt.Errorf("unable to create destination folder: %s", err.Error())
	}

	content, err := yaml.Marshal(tenant)
	if err != nil {
		return fmt.Errorf("unable to marshal yaml instances file: %s", err.Error())
	}

	if err := os.WriteFile(destFile, content, fsRights); err != nil {
		return fmt.Errorf("failed to write tenant file: %s", err.Error())
	}

	command.Println(fmt.Sprintf("tenant file successfully initialized. Please check %s file", destFile))

	return nil
}

func (cmd *TenantCmd) preRunE(command *cobra.Command, args []string) error {
	cmd.Command.SilenceUsage = true
	dest, err := processDestination(cmd.Flags.Destination)
	if err == nil {
		cmd.Flags.Destination = dest
		return nil
	}

	return err
}
