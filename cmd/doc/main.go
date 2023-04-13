// Binary bbsctl is a command line interface to manage BigBlueSwarm load balancer
package main

import (
	"github.com/bigblueswarm/bbsctl/pkg/cmd/root"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := root.NewCmd()
	cmd.DisableAutoGenTag = true
	if err := doc.GenMarkdownTree(cmd, "./docs"); err != nil {
		panic(err)
	}
}
