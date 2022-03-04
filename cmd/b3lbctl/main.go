package main

import (
	"os"

	"github.com/SLedunois/b3lbctl/pkg/cmd/root"
	"github.com/SLedunois/b3lbctl/pkg/system"
)

func main() {
	if err := root.NewCmd().Execute(); err != nil {
		os.Exit(system.OperationNotPermittedExitCode)
	}
}
