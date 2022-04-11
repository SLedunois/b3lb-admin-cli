package get

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCmd(t *testing.T) {
	t.Run("Launching instances command should not return any error", func(t *testing.T) {
		err := NewCmd().Execute()
		assert.Nil(t, err)
	})
}
