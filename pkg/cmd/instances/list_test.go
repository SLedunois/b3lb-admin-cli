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
	"github.com/SLedunois/b3lbctl/internal/test"
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

	tests := []test.CmdTest{
		{
			Name: "an error thrown by admin should return an error",
			Mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return []api.BigBlueButtonInstance{}, errors.New("admin error")
				}
			},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "calling list cmd with --json should print result as a json response",
			Mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return instances, nil
				}
			},
			Args: []string{"--json"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				assert.Equal(t, text.NewJSONTransformer("", "  ")(instances), strings.TrimSpace(string(out)))
			},
		},
		{
			Name: "calling list cmd with --csv should print result as a csv result",
			Mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return instances, nil
				}
			},
			Args: []string{"--csv"},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := fmt.Sprintf(`Url,Secret
%s,%s`, url, secret)
				assert.Equal(t, expected, strings.TrimSpace(string(out)))
			},
		},
		{
			Name: "calling list cmd with no flag should return a formatted table",
			Mock: func() {
				mock.ListAdminFunc = func() ([]api.BigBlueButtonInstance, error) {
					return instances, nil
				}
			},
			Args: []string{},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := fmt.Sprintf(`Url                             Secret  
%s  %s`, url, secret)
				assert.Equal(t, expected, strings.TrimSpace(string(out)))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewListCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			test.Mock()
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
