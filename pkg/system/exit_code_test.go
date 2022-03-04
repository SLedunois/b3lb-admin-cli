package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExitCode(t *testing.T) {
	assert.Equal(t, 1, OperationNotPermittedExitCode)
	assert.Equal(t, 2, NoSuchFileOrDirectoryExitCode)
	assert.Equal(t, 11, ResourceTemporarilyUnavailableExitCode)
}
