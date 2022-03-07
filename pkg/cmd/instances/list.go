package instances

import (
	"fmt"

	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

var jsonOutput bool
var csvOutput bool

// NewListCmd return the instances list subcommand
func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List BigBlueButton instances",
		Long:  `List all BigBlueButton instances available in your B3LB cluster`,
		RunE:  list,
	}

	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "json output")
	cmd.Flags().BoolVarP(&csvOutput, "csv", "c", false, "csv output")

	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	instances, err := admin.API.List()
	if err != nil {
		return fmt.Errorf("an error occured when getting remote instances: %s", err.Error())
	}

	if jsonOutput {
		renderJSON(cmd, instances)
	} else if csvOutput {
		renderTable(cmd, instances, true)
	} else {
		renderTable(cmd, instances, false)
	}

	return nil
}

func renderTable(cmd *cobra.Command, instances []api.BigBlueButtonInstance, csv bool) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"URL", "Secret"})

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
