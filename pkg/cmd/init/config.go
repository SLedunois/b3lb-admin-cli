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

type InitConfigCmd struct {
	Command *cobra.Command
	Flags   *InitConfigFlags
}

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

	return cmd.Command
}

func (cmd *InitConfigCmd) ApplyFlags() {
	cmd.Command.Flags().StringVarP(&cmd.Flags.B3LB, "b3lb", "b", "", "B3lb url")
	cmd.Command.Flags().StringVarP(&cmd.Flags.APIKey, "key", "k", "", "B3lb admin api key")
	cmd.Command.Flags().StringVarP(&cmd.Flags.Destination, "destination", "d", b3lbconfig.DefaultConfigFolder, "Configuration file folder destination")
}

func (cmd *InitConfigCmd) init(command *cobra.Command, args []string) error {
	cmd.Command.SilenceUsage = true
	dest := cmd.Flags.Destination
	if dest == b3lbconfig.DefaultConfigFolder {
		formalized, err := formalizeDefaultConfigFolder()
		if err != nil {
			return fmt.Errorf("unable to initialize configuration: %s", err.Error())
		}

		dest = formalized
	}

	filePath := path.Join(dest, config.DefaultConfigFileName)

	if info, _ := os.Stat(filePath); info != nil {
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
		return fmt.Errorf("unable to marshal yaml configuation: %s", err.Error())
	}

	if err := os.WriteFile(path.Join(dest, config.DefaultConfigFileName), content, fsRights); err != nil {
		return fmt.Errorf("failed to write configuration file: %s", err.Error())
	}

	return nil
}

func formalizeDefaultConfigFolder() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".b3lb"), nil
}
