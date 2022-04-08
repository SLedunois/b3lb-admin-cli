package init

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCmd(t *testing.T) {
	t.Run("Launch init command should not return any error", func(t *testing.T) {
		assert.Nil(t, NewCmd().Execute())
	})
}
