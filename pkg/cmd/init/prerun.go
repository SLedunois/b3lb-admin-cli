package init

import (
	"fmt"
	"os"
	"path"

	b3lbconfig "github.com/SLedunois/b3lb/v2/pkg/config"
)

func processDestination(dest string) (string, error) {
	if dest == b3lbconfig.DefaultConfigFolder {
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

	return path.Join(home, ".b3lb"), nil
}
