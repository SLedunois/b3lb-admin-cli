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

const (
	tenantFileNameFormatter = "%s.tenant.yml"

	initTenantExample = `
bbsctl init tenant --host bbs.example.com
# generates the following file
#
# kind: Tenant
# spec:
    host: bbs.example.com
# instances: []

bbsctl init tenant --host bbs.example.com --dest /path/to/file

bbsctl init tenant --host bbs.example.com --meeting_pool 10
# generates the following file
#
# kind: Tenant
# spec:
#    host: bbs.example.com
#    meeting_pool: 10
# instances: []

bbsctl init tenant --host bbs.example.com --meeting_pool 10 --user_pool 100
# generates the following file
#
# kind: Tenant
# spec:
#    host: bbs.example.com
#    meeting_pool: 10
#    user_pool: 100
# instances: []
	`
)

// TenantCmd represents the `bbsctl init tenant` command
type TenantCmd struct {
	Command *cobra.Command
	Flags   *TenantFlags
}

// NewInitTenantCmd initialize a new TenantCmd
func NewInitTenantCmd() *cobra.Command {
	cmd := &TenantCmd{
		Command: &cobra.Command{
			Use:     "tenant [flags]",
			Short:   "Initialize a new bigblueswarm tenant configuration file",
			Long:    "Initialize a new bigblueswarm tenant configuration file if not exits",
			Example: initTenantExample,
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
	cmd.Command.Flags().Int64VarP(&cmd.Flags.MeetingPool, "meeting_pool", "", -1, "Tenant meeting pool. This means the tenant can't create more meetings than the configured meeting pool. -1 is ignored.")
	cmd.Command.Flags().Int64VarP(&cmd.Flags.UserPool, "user_pool", "", -1, "Tenant user pool. This means the tenant can't have more users than the configured user pool. -1 is ignored.")
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

	if cmd.Flags.MeetingPool != -1 {
		tenant.Spec.MeetingsPool = &cmd.Flags.MeetingPool
	}

	if cmd.Flags.UserPool != -1 {
		tenant.Spec.UserPool = &cmd.Flags.UserPool
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
