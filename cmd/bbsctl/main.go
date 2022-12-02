package main

import (
	"os"

	"github.com/bigblueswarm/bbsctl/pkg/cmd/root"
	"github.com/bigblueswarm/bbsctl/pkg/system"
)

func main() {
	if err := root.NewCmd().Execute(); err != nil {
		os.Exit(system.OperationNotPermittedExitCode)
	}
}
