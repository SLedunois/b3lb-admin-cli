package test

import (
	"bytes"
	"testing"
)

// CmdTest struct is a table driven test object struct
type CmdTest struct {
	Name      string
	Mock      func()
	Args      []string
	Validator func(t *testing.T, output *bytes.Buffer, err error)
}
