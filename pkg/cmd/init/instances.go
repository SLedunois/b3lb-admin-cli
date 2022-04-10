package init

import (
	"fmt"
	"os"
	"path"

	b3lbadmin "github.com/SLedunois/b3lb/v2/pkg/admin"
	b3lbconfig "github.com/SLedunois/b3lb/v2/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const instancesConfigFilename = "instances.yml"

// InstancesCmd represents the `b3lbctl init instances` command
type InstancesCmd struct {
	Command *cobra.Command
	Flags   *InstancesFlags
}

// NewInitInstancesCmd initialize a new InstancesCmd
func NewInitInstancesCmd() *cobra.Command {
	cmd := &InstancesCmd{
		Command: &cobra.Command{
			Use:   "instances [flags]",
			Short: "Initialize b3lb instances file",
			Long:  "Create instances list file if it does not exists",
		},
		Flags: NewInitInstancesFlags(),
	}

	cmd.ApplyFlags()

	cmd.Command.RunE = cmd.init
	cmd.Command.PreRunE = cmd.preRunE

	return cmd.Command
}

// ApplyFlags apply command flags to InitInstancesCmd
func (cmd *InstancesCmd) ApplyFlags() {
	cmd.Command.Flags().StringVarP(&cmd.Flags.Destination, "dest", "d", b3lbconfig.DefaultConfigFolder, "File folder destination")
}

func (cmd *InstancesCmd) init(command *cobra.Command, args []string) error {
	destFile := path.Join(cmd.Flags.Destination, instancesConfigFilename)
	if fileAlreadyExists(destFile) {
		return fmt.Errorf("instances configuration file already exists. Please consider editing %s file", destFile)
	}

	instances := &b3lbadmin.InstanceList{
		Kind: "InstanceList",
	}

	if err := os.MkdirAll(cmd.Flags.Destination, fsRights); err != nil {
		return fmt.Errorf("unable to create destination folder: %s", err.Error())
	}

	content, err := yaml.Marshal(instances)
	if err != nil {
		return fmt.Errorf("unable to marshal yaml instances file: %s", err.Error())
	}

	if err := os.WriteFile(destFile, content, fsRights); err != nil {
		return fmt.Errorf("failed to write instances file: %s", err.Error())
	}

	command.Println(fmt.Sprintf("instances file successfully initialized. Please check %s file", destFile))

	return nil
}

func (cmd *InstancesCmd) preRunE(command *cobra.Command, args []string) error {
	cmd.Command.SilenceUsage = true
	dest, err := processDestination(cmd.Flags.Destination)
	if err == nil {
		cmd.Flags.Destination = dest
		return nil
	}

	return err
}
