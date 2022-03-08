package instances

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	type test struct {
		name      string
		args      []string
		mock      func()
		validator func(t *testing.T, output *bytes.Buffer, err error)
	}

	mock.InitAdminMock()

	tests := []test{
		{
			name: "no url flag should return a cmd error",
			args: []string{},
			mock: func() {},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an error returned by admin api should return a cmd error and log a valid message",
			args: []string{"--url", "http://localhost"},
			mock: func() {
				mock.DeleteAdminFunc = func(url string) error {
					return errors.New("admin error")
				}
			},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to delete http://localhost BigBlueButton instance from your cluster: admin error", err.Error())
			},
		},
		{
			name: `a valid command should log "instance deleted" message and return no error`,
			args: []string{"--url", "http://localhost"},
			mock: func() {
				mock.DeleteAdminFunc = func(url string) error {
					return nil
				}
			},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				assert.Equal(t, "instance http://localhost deleted\n", string(out))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			b := bytes.NewBufferString("")
			cmd := NewDeleteCmd()
			cmd.SetArgs(test.args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.validator(t, b, err)
		})
	}
}
