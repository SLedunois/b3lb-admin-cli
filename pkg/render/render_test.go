package render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableStyle(t *testing.T) {
	assert.NotNil(t, TableStyle())
}
