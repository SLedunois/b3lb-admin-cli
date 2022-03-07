package instances

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

func TestListCmd(t *testing.T) {
	url := "http://localhost/bigbluebutton"
	secret := "secret"

	instances := []api.BigBlueButtonInstance{
		{
			URL:    url,
			Secret: secret,
		},
	}

	mock.InitAdminMock()

	type test struct {
		name      string
		mock      func()
		args      []string
		validator func(t *testing.T, output *bytes.Buffer, err error)
	}

	tests := []test{
		{
			name: "an error thrown by admin should return an error",
			mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return []api.BigBlueButtonInstance{}, errors.New("admin error")
				}
			},
			args: []string{},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "calling list cmd with --json should print result as a json response",
			mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return instances, nil
				}
			},
			args: []string{"--json"},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				assert.Equal(t, text.NewJSONTransformer("", "  ")(instances), strings.TrimSpace(string(out)))
			},
		},
		{
			name: "calling list cmd with --csv should print result as a csv result",
			mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return instances, nil
				}
			},
			args: []string{"--csv"},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := fmt.Sprintf(`URL,Secret
%s,%s`, url, secret)
				assert.Equal(t, expected, strings.TrimSpace(string(out)))
			},
		},
		{
			name: "calling list cmd with no flag should return a formatted table",
			mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return instances, nil
				}
			},
			args: []string{},
			validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := fmt.Sprintf(`URL                             Secret  
%s  %s`, url, secret)
				assert.Equal(t, expected, strings.TrimSpace(string(out)))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewListCmd()
			cmd.SetArgs(test.args)
			cmd.SetOut(b)
			test.mock()
			err := cmd.Execute()
			test.validator(t, b, err)
		})
	}
}
