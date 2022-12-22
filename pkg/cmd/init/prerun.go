// Package init provides the init command
package init

import (
	"fmt"
	"os"
	"path"

	bbsconfig "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
)

func processDestination(dest string) (string, error) {
	if dest == bbsconfig.DefaultConfigFolder {
		formalized, err := formalizeDefaultConfigFolder()
		if err != nil {
			return "", fmt.Errorf("unable to initialize configuration: %s", err.Error())
		}

		dest = formalized
	}

	return dest, nil
}

func formalizeDefaultConfigFolder() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".bigblueswarm"), nil
}
