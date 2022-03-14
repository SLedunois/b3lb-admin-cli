package clusterinfo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"

	"github.com/SLedunois/b3lb/pkg/balancer"
	"github.com/SLedunois/b3lbctl/internal/mock"
	"github.com/SLedunois/b3lbctl/internal/test"
)

func TestClusterInfo(t *testing.T) {
	mock.InitAdminMock()
	tests := []test.CmdTest{
		{
			Name: "an error thrown by admin cluster status method should return an error",
			Mock: func() {
				mock.ClusterStatusAdminFunc = func() ([]balancer.InstanceStatus, error) {
					return nil, errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an error thrown by admin b3lb api status method should return an error",
			Mock: func() {
				mock.ClusterStatusAdminFunc = func() ([]balancer.InstanceStatus, error) {
					return []balancer.InstanceStatus{}, nil
				}
				mock.B3lbAPIStatusAdminFunc = func() (string, error) {
					return "", errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "a valid command should print cluster status",
			Mock: func() {
				mock.ClusterStatusAdminFunc = func() ([]balancer.InstanceStatus, error) {
					return []balancer.InstanceStatus{
						{
							Host:               "http://localhost/bigbluebutton",
							CPU:                8.32,
							Mem:                55.36,
							ActiveMeeting:      1,
							ActiveParticipants: 10,
							APIStatus:          "Up",
						},
					}, nil
				}
				mock.B3lbAPIStatusAdminFunc = func() (string, error) {
					return "Up", nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				out, outErr := ioutil.ReadAll(output)
				assert.Nil(t, outErr)
				expected := `B3LB API             Up  
Active meetings      1   
Active participants  10  

API  Host                            CPU     Mem      Active Meetings  Active Participants  
Up   http://localhost/bigbluebutton  8.32 %  55.36 %                1                   10`
				assert.Equal(t, expected, strings.TrimSpace(string(out)))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			text.DisableColors() // Disable colors to compare printed result
			cmd := NewCmd()
			cmd.SetArgs([]string{})
			cmd.SetOut(b)
			test.Mock()
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
