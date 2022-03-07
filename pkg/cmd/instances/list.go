package instances

import (
	"fmt"

	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// ListCmd struct represents the list command object
type ListCmd struct {
	Command *cobra.Command
	Flags   *ListFlags
}

// NewListCmd return the instances list subcommand
func NewListCmd() *cobra.Command {
	cmd := &ListCmd{
		Command: &cobra.Command{
			Use:   "list",
			Short: "List BigBlueButton instances",
			Long:  `List all BigBlueButton instances available in your B3LB cluster`,
		},
		Flags: NewListFlags(),
	}

	cmd.Command.RunE = cmd.list

	cmd.ApplyFlags()

	return cmd.Command
}

func (cmd *ListCmd) list(command *cobra.Command, args []string) error {
	instances, err := admin.API.List()
	if err != nil {
		return fmt.Errorf("an error occured when getting remote instances: %s", err.Error())
	}

	if cmd.Flags.JSON {
		renderJSON(command, instances)
	} else if cmd.Flags.CSV {
		renderTable(command, instances, true)
	} else {
		renderTable(command, instances, false)
	}

	return nil
}

func renderTable(cmd *cobra.Command, instances []api.BigBlueButtonInstance, csv bool) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Url", "Secret"})

	for _, instance := range instances {
		t.AppendRow(table.Row{instance.URL, instance.Secret})
	}

	t.SetStyle(tableStyle())
	if csv {
		cmd.Println(t.RenderCSV())
	} else {
		cmd.Println(t.Render())
	}
}

func renderJSON(cmd *cobra.Command, instances []api.BigBlueButtonInstance) {
	cmd.Println(text.NewJSONTransformer("", "  ")(instances))
}

func tableStyle() table.Style {
	return table.Style{
		Name: "Docker style",
		Box: table.BoxStyle{
			BottomLeft:       "",
			BottomRight:      "",
			BottomSeparator:  "",
			Left:             "",
			LeftSeparator:    "",
			MiddleHorizontal: "",
			MiddleSeparator:  "",
			MiddleVertical:   "",
			PaddingLeft:      "",
			PaddingRight:     "  ",
			Right:            "",
			RightSeparator:   "",
			TopLeft:          "",
			TopRight:         "",
			TopSeparator:     "",
			UnfinishedRow:    "",
		},
		Format: table.FormatOptions{
			Header: text.FormatTitle,
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
	}
}
