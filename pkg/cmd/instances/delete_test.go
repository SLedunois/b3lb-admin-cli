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

func TestDelete(t *testing.T) {

	mock.InitAdminMock()

	tests := []test.CmdTest{
		{
			Name: "no url flag should return a cmd error",
			Args: []string{},
			Mock: func() {},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an error returned by admin api should return a cmd error and log a valid message",
			Args: []string{"--url", "http://localhost"},
			Mock: func() {
				mock.DeleteAdminFunc = func(url string) error {
					return errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to delete http://localhost BigBlueButton instance from your cluster: admin error", err.Error())
			},
		},
		{
			Name: `a valid command should log "instance deleted" message and return no error`,
			Args: []string{"--url", "http://localhost"},
			Mock: func() {
				mock.DeleteAdminFunc = func(url string) error {
					return nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				assert.Equal(t, "instance http://localhost deleted\n", string(out))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mock()
			b := bytes.NewBufferString("")
			cmd := NewDeleteCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
