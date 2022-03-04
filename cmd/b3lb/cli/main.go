package main

import (
	"fmt"
	"os"

	"github.com/SLedunois/b3lb-admin-cli/pkg/cmd/root"
	"github.com/SLedunois/b3lb-admin-cli/pkg/system"
)

func main() {
	if err := root.NewCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(system.OperationNotPermittedExitCode)
	}
}
