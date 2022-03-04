package instances

import (
	"fmt"

	"github.com/SLedunois/b3lbctl/pkg/admin"
	"github.com/spf13/cobra"
)

var url string
var secret string

// NewAddCmd return the instances add subcommand
func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a BigBlueButton instance",
		Long:  `Add a BigBlueButton instance in your B3LB cluster`,
		RunE:  add,
	}

	cmd.Flags().StringVarP(&url, "url", "u", "", "BigBlueButton instance url")
	cmd.MarkFlagRequired("url")
	cmd.Flags().StringVarP(&secret, "secret", "s", "", "BigBlueButton secret")
	cmd.MarkFlagRequired("secret")

	return cmd
}

func add(cmd *cobra.Command, args []string) error {
	if err := admin.API.Add(url, secret); err != nil {
		cmd.SilenceUsage = true
		return fmt.Errorf("unable to add BigBlueButton instance to your cluster: %s", err.Error())
	}

	cmd.Println(fmt.Sprintf(`instance %s added`, url))
	return nil
}
