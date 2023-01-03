// Package init provides the init command
package init

import (
	"fmt"
	"os"
	"path"

	"github.com/bigblueswarm/bbsctl/pkg/config"
	bbsconfig "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const (
	fsRights             = 0755
	initConfigCmdExample = `
bbsctl init config --dest /path/to/files

bbsctl init config
# generates the following file
#
# bbs: ""
# apiKey: ""

bbsctl init config --bbs http://bbs.example.com
# generates the following file
# 
# bbs: http://bbs.example.com
# apiKey: ""

bbsctl init config --bbs http://bbs.example.com --key api_key
# generates the following file
#
# bbs: http://bbs.example.com
# apiKey: api_key
`
)

// InitConfigCmd represents the `bbsctl init config` command
type InitConfigCmd struct {
	Command *cobra.Command
	Flags   *ConfigFlags
}

// NewInitConfigCmd initialize a new InitConfigCmd struct
func NewInitConfigCmd() *cobra.Command {
	cmd := &InitConfigCmd{
		Command: &cobra.Command{
			Use:     "config [flags]",
			Short:   "Initialize bbsctl configuration",
			Long:    "Create bbsctl if not exists and initialize a basic configuration",
			Example: initConfigCmdExample,
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
	cmd.Command.Flags().StringVarP(&cmd.Flags.BBS, "bbs", "b", "", "BigBlueSwarm url")
	cmd.Command.Flags().StringVarP(&cmd.Flags.APIKey, "key", "k", "", "BigBlueSwarm admin api key")
	cmd.Command.Flags().StringVarP(&cmd.Flags.Destination, "dest", "d", bbsconfig.DefaultConfigFolder, "Configuration file folder destination")
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
		BBS:    cmd.Flags.BBS,
		APIKey: cmd.Flags.APIKey,
	}

	content, err := yaml.Marshal(conf)
	if err != nil {
		return fmt.Errorf("unable to marshal yaml configuration: %s", err.Error())
	}

	if err := os.WriteFile(path.Join(dest, config.DefaultConfigFileName), content, fsRights); err != nil {
		return fmt.Errorf("failed to write configuration file: %s", err.Error())
	}

	command.Println(fmt.Sprintf("configuration successfully initialized. Please check %s file", path.Join(dest, config.DefaultConfigFileName)))

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
