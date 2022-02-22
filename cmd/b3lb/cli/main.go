package main

import (
	"fmt"
	"os"

	"github.com/SLedunois/b3lb-admin-cli/pkg/cmd/root"
)

func main() {
	if err := root.NewCmdRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	  }
}
