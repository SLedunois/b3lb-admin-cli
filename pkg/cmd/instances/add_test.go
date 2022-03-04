package instances

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestAddCmd(t *testing.T) {
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
			args: []string{"--secret", "dummy_secret"},
			mock: func() {},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "no secret flag should return a cmd error",
			args: []string{"--url", "http://localhost"},
			mock: func() {},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an error returned by admin api should return a cmd error and log a valid message",
			args: []string{"--url", "http://localhost", "--secret", "dummy_secret"},
			mock: func() {
				mock.AddAdminFunc = func(url, secret string) error {
					return errors.New("admin error")
				}
			},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to add BigBlueButton instance to your cluster: admin error", err.Error())
			},
		},
		{
			name: `a valid command should log "instance added" message and return no error`,
			args: []string{"--url", "http://localhost", "--secret", "dummy_secret"},
			mock: func() {
				mock.AddAdminFunc = func(url, secret string) error {
					return nil
				}
			},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				assert.Equal(t, "instance http://localhost added\n", string(out))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			b := bytes.NewBufferString("")
			cmd := NewAddCmd()
			cmd.SetArgs(test.args)
			cmd.SetOut(b)
			err := cmd.Execute()
			test.validator(t, b, err)
		})
	}
}
