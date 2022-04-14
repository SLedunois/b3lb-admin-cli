package delete

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCmd(t *testing.T) {
	assert.NotNil(t, NewCmd())
}
