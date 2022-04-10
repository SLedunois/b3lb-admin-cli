package init

import (
	"fmt"
	"os"
	"path"

	b3lbconfig "github.com/SLedunois/b3lb/v2/pkg/config"
	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const fsRights = 0755

// InitConfigCmd represents the `b3lbctl init config` command
type InitConfigCmd struct {
	Command *cobra.Command
	Flags   *ConfigFlags
}

// NewInitConfigCmd initialize a new InitConfigCmd struct
func NewInitConfigCmd() *cobra.Command {
	cmd := &InitConfigCmd{
		Command: &cobra.Command{
			Use:   "config [flags]",
			Short: "Initialize b3lbctl configuration",
			Long:  "Create b3lbctl if not exists and initialize a basic configuration",
		},
		Flags: NewInitConfigFlags(),
	}

	cmd.ApplyFlags()

	cmd.Command.RunE = cmd.init
	cmd.Command.PreRunE = cmd.prerun

	return cmd.Command
}

// ApplyFlags apply command flags to InitConfigCmd
func (cmd *InitConfigCmd) ApplyFlags() {
	cmd.Command.Flags().StringVarP(&cmd.Flags.B3LB, "b3lb", "b", "", "B3lb url")
	cmd.Command.Flags().StringVarP(&cmd.Flags.APIKey, "key", "k", "", "B3lb admin api key")
	cmd.Command.Flags().StringVarP(&cmd.Flags.Destination, "dest", "d", b3lbconfig.DefaultConfigFolder, "Configuration file folder destination")
}

func (cmd *InitConfigCmd) init(command *cobra.Command, args []string) error {
	dest := cmd.Flags.Destination
	filePath := path.Join(dest, config.DefaultConfigFileName)

	if fileAlreadyExists(filePath) {
		return fmt.Errorf("configuration already exists, see %s", filePath)
	}

	if err := os.MkdirAll(dest, fsRights); err != nil {
		return fmt.Errorf("unable to create destination folder: %s", err.Error())
	}

	conf := &config.Config{
		B3lb:   cmd.Flags.B3LB,
		APIKey: cmd.Flags.APIKey,
	}

	content, err := yaml.Marshal(conf)
	if err != nil {
		return fmt.Errorf("unable to marshal yaml configuration: %s", err.Error())
	}

	if err := os.WriteFile(path.Join(dest, config.DefaultConfigFileName), content, fsRights); err != nil {
		return fmt.Errorf("failed to write configuration file: %s", err.Error())
	}

	return nil
}

func (cmd *InitConfigCmd) prerun(command *cobra.Command, args []string) error {
	cmd.Command.SilenceUsage = true
	dest, err := processDestination(cmd.Flags.Destination)
	if err == nil {
		cmd.Flags.Destination = dest
		return nil
	}

	return err
}
