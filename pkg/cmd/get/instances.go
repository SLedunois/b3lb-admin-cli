package get

import (
	"fmt"

	"github.com/bigblueswarm/bbsctl/pkg/admin"
	"github.com/bigblueswarm/bbsctl/pkg/render"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/api"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// InstancesCmd struct represents the list command object
type InstancesCmd struct {
	Command *cobra.Command
	Flags   *Flags
}

// NewInstancesCmd return the instances list subcommand
func NewInstancesCmd() *cobra.Command {
	cmd := &InstancesCmd{
		Command: &cobra.Command{
			Use:   "instances [flags]",
			Short: "Display all BigBlueButton instances available in your BigBlueSwarm cluster",
			Long:  `Display all BigBlueButton instances available in your BigBlueSwarm cluster`,
		},
		Flags: NewFlags(),
	}

	cmd.Command.RunE = cmd.list

	cmd.ApplyFlags()

	return cmd.Command
}

// ApplyFlags apply GetFlags to provided command
func (cmd *InstancesCmd) ApplyFlags() {
	cmd.Command.Flags().BoolVarP(&cmd.Flags.CSV, "csv", "c", cmd.Flags.CSV, "csv output")
	cmd.Command.Flags().BoolVarP(&cmd.Flags.JSON, "json", "j", cmd.Flags.JSON, "json output")
}

func (cmd *InstancesCmd) list(command *cobra.Command, args []string) error {
	instances, err := admin.API.List()
	if err != nil {
		return fmt.Errorf("an error occured when getting remote instances: %s", err.Error())
	}

	if cmd.Flags.JSON {
		renderBigBlueButtonInstancesJSON(command, instances)
	} else {
		renderBigBlueButtonInstancesTable(command, instances, cmd.Flags.CSV)
	}

	return nil
}

func renderBigBlueButtonInstancesTable(cmd *cobra.Command, instances []api.BigBlueButtonInstance, csv bool) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Url", "Secret"})

	for _, instance := range instances {
		t.AppendRow(table.Row{instance.URL, instance.Secret})
	}

	t.SetStyle(render.TableStyle())
	if csv {
		cmd.Println(t.RenderCSV())
	} else {
		cmd.Println(t.Render())
	}
}

func renderBigBlueButtonInstancesJSON(cmd *cobra.Command, instances []api.BigBlueButtonInstance) {
	cmd.Println(text.NewJSONTransformer("", "  ")(instances))
}
