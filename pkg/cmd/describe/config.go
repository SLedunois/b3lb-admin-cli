// Package describe provides the describe command
package describe

import (
	"fmt"
	"strings"

	"github.com/bigblueswarm/bbsctl/pkg/admin"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// ConfigCmd represents config command type
type ConfigCmd struct {
	Flags   *ConfigFlags
	Command *cobra.Command
}

// NewConfigCmd initialize a new view command
func NewConfigCmd() *cobra.Command {
	cmd := &ConfigCmd{
		Flags: NewConfigFlags(),
		Command: &cobra.Command{
			Use:   "config",
			Short: "describe BigBlueSwarm configuration.",
			Long:  `describe BigBlueSwarm configuration.`,
		},
	}

	cmd.Command.RunE = cmd.process

	return cmd.Command
}

func (c *ConfigCmd) process(command *cobra.Command, args []string) error {
	command.SilenceUsage = true
	config, err := admin.API.GetConfiguration()
	if err != nil {
		return fmt.Errorf("unable to describe bigblueswarm configuration: %s", err.Error())
	}

	out, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("unable to describe bigblueswarm configuration: %s", err.Error())
	}

	command.Println(normalizeYaml(string(out)))

	return nil
}

func normalizeYaml(value string) string {
	value = strings.ReplaceAll(value, "---", "")
	value = strings.ReplaceAll(value, "...", "")
	return strings.TrimSpace(value)
}
