package main

import (
	"fmt"
	"os"

	"github.com/SLedunois/b3lbctl/pkg/cmd/root"
	"github.com/SLedunois/b3lbctl/pkg/system"
)

func main() {
	if err := root.NewCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(system.OperationNotPermittedExitCode)
	}
}
