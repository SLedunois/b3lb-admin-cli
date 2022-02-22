package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	return &cobra.Command{
		Use:   "b3lb-admin <command> [flags]",
		Short: "B3LB admin cli",
		Long:  `Manage your B3LB cluster from the command line`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("In b3lb-admin command")
		},
	}
}
