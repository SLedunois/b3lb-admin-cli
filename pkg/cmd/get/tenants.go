package get

import (
	"fmt"

	b3lbadmin "github.com/SLedunois/b3lb/v2/pkg/admin"
	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/SLedunois/b3lbctl/pkg/render"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

type TenantsCmd struct {
	Command *cobra.Command
	Flags   *Flags
}

// NewTenantCmd returns the get tenants subcommand
func NewTenantCmd() *cobra.Command {
	cmd := &TenantsCmd{
		Command: &cobra.Command{
			Use:   "tenants [flags]",
			Short: "Display all B3lb tenants available in your B3LB cluster",
			Long:  "Display all B3lb tenants available in your B3LB cluster",
		},
		Flags: NewFlags(),
	}

	cmd.ApplyFlags()

	cmd.Command.RunE = cmd.list

	return cmd.Command
}

// ApplyFlags apply ListFlags to provided command
func (cmd *TenantsCmd) ApplyFlags() {
	cmd.Command.Flags().BoolVarP(&cmd.Flags.CSV, "csv", "c", cmd.Flags.CSV, "csv output")
	cmd.Command.Flags().BoolVarP(&cmd.Flags.JSON, "json", "j", cmd.Flags.JSON, "json output")
}

func (cmd *TenantsCmd) list(command *cobra.Command, args []string) error {
	tenants, err := admin.API.GetTenants()
	if err != nil {
		return fmt.Errorf("unable to fetch tenants: %s", err.Error())
	}

	if cmd.Flags.JSON {
		command.Println(text.NewJSONTransformer("", "  ")(tenants.Tenants))
	} else {
		renderTenantsTable(command, tenants, cmd.Flags.CSV)
	}

	return nil
}

func renderTenantsTable(cmd *cobra.Command, tenants *b3lbadmin.TenantList, csv bool) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Hostname", "Instances"})

	for _, tenant := range tenants.Tenants {
		t.AppendRow(table.Row{tenant.Hostname, tenant.InstanceCount})
	}

	t.SetStyle(render.TableStyle())
	if csv {
		cmd.Println(t.RenderCSV())
	} else {
		cmd.Println(t.Render())
	}
}
