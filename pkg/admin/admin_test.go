package admin

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lb/pkg/restclient"
	restmock "github.com/SLedunois/b3lb/pkg/restclient/mock"
	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/stretchr/testify/assert"
)

var apiKey = "api_key"
var instance = "https://localhost:8090"

type test struct {
	name      string
	mock      func()
	validator func(t *testing.T, instances []api.BigBlueButtonInstance, err error)
}

func initTests() {
	Init()
	config.APIKey = &apiKey
	config.URL = &instance
	restclient.Client = &restmock.RestClient{}

}

func TestList(t *testing.T) {
	initTests()
	instances := []api.BigBlueButtonInstance{
		{
			URL:    "http://localhost/bigbluebutton",
			Secret: "dummy_secret",
		},
	}

	tests := []test{
		{
			name: "an error thrown by restclient should return an error",
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			instances, err := API.List()
			test.validator(t, instances, err)
		})
	}
}

func TestAdd(t *testing.T) {
	initTests()
	tests := []test{
		{
			name: "add method should return an error if restclient return one",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{}, errors.New("http error")
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "add method should not return an error if restclient return a valid response",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusCreated,
					}, nil
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "add method should return an error if restclient does not return a 201 - HTTP Created - status response",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := API.Add("http://localhost:8080/bigbluebutton", "secret")
			test.validator(t, nil, err)
		})
	}
}

func TestDelete(t *testing.T) {
	initTests()
	tests := []test{
		{
			name: "an error returned by the restclient should return an error",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, errors.New("http error")
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "no error returned by the restclient but an http code != 204 should return an error",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "no error and a 204 http status should return no error",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNoContent,
					}, nil
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "no error and a 404 http status should return an error",
			mock: func() {
				restmock.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNotFound,
					}, nil
				}
			},
			validator: func(t *testing.T, instances []api.BigBlueButtonInstance, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := API.Delete("http://localhost:8080/bigbluebutton")
			test.validator(t, nil, err)
		})
	}
}
