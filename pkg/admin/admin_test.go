package admin

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lb/pkg/restclient"
	restmock "github.com/SLedunois/b3lb/pkg/restclient/mock"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	apiKey := "api_key"
	url := "https://localhost:8090"

	instances := []api.BigBlueButtonInstance{
		{
			URL:    "http://localhost/bigbluebutton",
			Secret: "dummy_secret",
		},
	}

	type test struct {
		name      string
		mock      func()
		validator func(t *testing.T, instances []api.BigBlueButtonInstance, err error)
	}

	tests := []test{
		{
			name: "an error thrown by restclien should return an error",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("http error")
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an error thrown by json unmarshaller should return an error",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(nil)),
					}, nil
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a valid request should a valid instances list",
			mock: func() {
				resp, err := json.Marshal(instances)
				if err != nil {
					panic(err)
				}

				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(resp)),
					}, nil
				}
			},
			validator: func(t *testing.T, bbb []api.BigBlueButtonInstance, err error) {
				assert.Nil(t, err)
				assert.Equal(t, instances, bbb)
			},
		},
	}

	Init()
	config.APIKey = &apiKey
	config.URL = &url
	restclient.Client = &restmock.RestClient{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			instances, err := API.List()
			test.validator(t, instances, err)
		})
	}
}
