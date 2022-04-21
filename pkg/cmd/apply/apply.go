package apply

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Cmd represents appl command type
type Cmd struct {
	Command *cobra.Command
	Flags   *Flags
}

// NewCmd initialize a new Apply command
func NewCmd() *cobra.Command {
	cmd := &Cmd{
		Command: &cobra.Command{
			Use:   "apply -f [filepath]",
			Short: "Apply a configuration to b3lb server using a file",
			Long:  "Apply a configuration to b3lb server using a file",
		},
		Flags: NewFlags(),
	}

	cmd.ApplyFlags()

	cmd.Command.RunE = cmd.process

	return cmd.Command
}

// ApplyFlags apply command flags to apply command
func (cmd *Cmd) ApplyFlags() {
	cmd.Command.Flags().StringVarP(&cmd.Flags.FilePath, "file", "f", "", "resource file path")
	cmd.Command.MarkFlagRequired("file")
}

func (cmd *Cmd) printApplyMessage(kind string, resource *interface{}) {
	if kind == "InstanceList" || kind == "Tenant" {
		cmd.Command.Println(fmt.Sprintf("%s resource created", kind))
	}
}

func (cmd *Cmd) process(command *cobra.Command, args []string) error {
	command.SilenceUsage = true
	b, err := cmd.loadFile()
	if err != nil {
		return fmt.Errorf("unable to apply file. File loading fail: %s", err.Error())
	}

	kind, resource, err := toResource(b)
	if err != nil {
		return fmt.Errorf("unable to apply file. File is not a valid resource: %s", err.Error())
	}

	if err := yaml.Unmarshal(b, &resource); err != nil {
		return fmt.Errorf("unable to apply file. Failed to unmarshal file content: %s", err.Error())
	}

	if err := admin.API.Apply(kind, &resource); err != nil {
		return fmt.Errorf("unable to apply file. %s", err.Error())
	}

	cmd.printApplyMessage(kind, &resource)

	return nil
}

func (cmd *Cmd) loadFile() ([]byte, error) {
	file, err := os.Open(cmd.Flags.FilePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			cmd.Command.Println(fmt.Sprintf("unable to close config file: %s", err))
		}
	}()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}
