package instances

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/SLedunois/b3lbctl/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestAddCmd(t *testing.T) {
	mock.InitAdminMock()

	tests := []test.CmdTest{
		{
			Name: "no url flag should return a cmd error",
			Args: []string{"--secret", "dummy_secret"},
			Mock: func() {},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "no secret flag should return a cmd error",
			Args: []string{"--url", "http://localhost"},
			Mock: func() {},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an error returned by admin api should return a cmd error and log a valid message",
			Args: []string{"--url", "http://localhost", "--secret", "dummy_secret"},
			Mock: func() {
				mock.AddAdminFunc = func(url, secret string) error {
					return errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to add BigBlueButton instance to your cluster: admin error", err.Error())
			},
		},
		{
			Name: `a valid command should log "instance added" message and return no error`,
			Args: []string{"--url", "http://localhost", "--secret", "dummy_secret"},
			Mock: func() {
				mock.AddAdminFunc = func(url, secret string) error {
					return nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				assert.Equal(t, "instance http://localhost added\n", string(out))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			b := bytes.NewBufferString("")
			cmd := NewAddCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
