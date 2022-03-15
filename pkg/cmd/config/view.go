package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/spf13/cobra"
)

// ViewCmd represents config command type
type ViewCmd struct {
	Flags   *ViewFlags
	Command *cobra.Command
}

// NewViewCmd initialize a new view command
func NewViewCmd() *cobra.Command {
	cmd := &ViewCmd{
		Flags: NewViewFlags(),
		Command: &cobra.Command{
			Use:   "view [flags]",
			Short: "Display B3LB config file.",
			Long:  `Display B3LB config file.`,
		},
	}

	cmd.Command.RunE = cmd.process

	return cmd.Command
}

func (c *ViewCmd) process(command *cobra.Command, args []string) error {
	file, err := os.Open(*config.Path)
	if err != nil {
		return err
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	command.Println(normalizeYaml(string(b)))

	return nil
}

func normalizeYaml(value string) string {
	value = strings.ReplaceAll(value, "---", "")
	value = strings.ReplaceAll(value, "...", "")
	return strings.TrimSpace(value)
}
