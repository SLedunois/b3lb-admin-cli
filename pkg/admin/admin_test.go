package admin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/SLedunois/b3lb/v2/pkg/admin"
	b3lbadmin "github.com/SLedunois/b3lb/v2/pkg/admin"
	"github.com/SLedunois/b3lb/v2/pkg/api"
	"github.com/SLedunois/b3lb/v2/pkg/balancer"
	b3lbconfig "github.com/SLedunois/b3lb/v2/pkg/config"
	"github.com/SLedunois/b3lb/v2/pkg/restclient"
	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/stretchr/testify/assert"
)

var apiKey = "api_key"
var instance = "https://localhost:8090"

type test struct {
	name      string
	mock      func()
	validator func(t *testing.T, value interface{}, err error)
}

func initTests() {
	Init()
	config.APIKey = &apiKey
	config.B3LB = &instance
	restclient.Client = &restclient.Mock{}

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
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("http error")
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an error thrown by json unmarshaller should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(nil)),
					}, nil
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
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

				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(resp)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, instances, value.([]api.BigBlueButtonInstance))
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
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{}, errors.New("http error")
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "add method should not return an error if restclient return a valid response",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusCreated,
					}, nil
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "add method should return an error if restclient does not return a 201 - HTTP Created - status response",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
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
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, errors.New("http error")
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "no error returned by the restclient but an http code != 204 should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "no error and a 204 http status should return no error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNoContent,
					}, nil
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "no error and a 404 http status should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNotFound,
					}, nil
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
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

func TestClusterStatus(t *testing.T) {
	initTests()
	bodyStatus := []balancer.InstanceStatus{
		{
			Host:               "http://localhost/bigbluebutton",
			CPU:                2.46,
			Mem:                66.34,
			ActiveMeeting:      0,
			ActiveParticipants: 0,
			APIStatus:          "Up",
		},
	}

	tests := []test{
		{
			name: "an error returned by restclient should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("rest error")
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an error thrown by json unmarshaller should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(nil)),
					}, nil
				}
			},
			validator: func(t *testing.T, _ interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a valid response should return a valid response",
			mock: func() {
				resp, err := json.Marshal(bodyStatus)
				if err != nil {
					panic(err)
				}

				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(resp)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, bodyStatus, value.([]balancer.InstanceStatus))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			status, err := API.ClusterStatus()
			test.validator(t, status, err)
		})
	}
}

func TestB3lbAPIStatus(t *testing.T) {
	initTests()
	check := api.CreateHealthCheck()

	tests := []test{
		{
			name: "an error returned by restclient should be returned and an empty string",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("rest error")
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "", value.(string))
			},
		},
		{
			name: "an error returned by xml unmarshaller should return an error and an empty string",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(nil)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "", value.(string))
			},
		},
		{
			name: `a failed status code should return a "Down" status`,
			mock: func() {
				check.ReturnCode = api.ReturnCodes().Failed
				value, err := xml.Marshal(check)
				if err != nil {
					panic(err)
				}

				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(value)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "Down", value.(string))
			},
		},
		{
			name: `a success status code should return a "Up" status`,
			mock: func() {
				check.ReturnCode = api.ReturnCodes().Success
				value, err := xml.Marshal(check)
				if err != nil {
					panic(err)
				}

				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(value)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "Up", value.(string))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			status, err := API.B3lbAPIStatus()
			test.validator(t, status, err)
		})
	}
}

func TestGetConfiguration(t *testing.T) {
	initTests()
	conf := &b3lbconfig.Config{
		Admin: b3lbconfig.AdminConfig{
			APIKey: "dummy_key",
		},
	}
	tests := []test{
		{
			name: "an error returned by restclient should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("rest error")
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an invalid body should return an error while unmarshalling content",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(nil)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a valid request should return a valid configuration",
			mock: func() {
				b, err := json.Marshal(conf)
				if err != nil {
					t.Fatal(err)
					return
				}

				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(b)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, value)
				conf := value.(*b3lbconfig.Config)
				assert.Nil(t, err)
				assert.Equal(t, "dummy_key", conf.Admin.APIKey)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			config, err := API.GetConfiguration()
			test.validator(t, config, err)
		})
	}
}

func TestGetTenants(t *testing.T) {
	initTests()
	tenants := &b3lbadmin.TenantList{
		Kind:    "TenantList",
		Tenants: []b3lbadmin.TenantListObject{},
	}

	tests := []test{
		{
			name: "an error returned by rest client should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("http error")
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a valid request should return a TenantList",
			mock: func() {
				value, err := json.Marshal(tenants)
				if err != nil {
					t.Fatal(err)
					return
				}

				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(value)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, tenants, value.(*b3lbadmin.TenantList))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			config, err := API.GetTenants()
			test.validator(t, config, err)
		})
	}
}

func TestGetTena(t *testing.T) {
	initTests()

	tenant := &b3lbadmin.Tenant{
		Kind: "Tenant",
		Spec: map[string]string{
			"host": "localhost",
		},
		Instances: []string{},
	}
	tests := []test{
		{
			name: "an error returned by restclient should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("rest error")
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a not found tenant should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNotFound,
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a b3lb internal server error shoul return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a valid request should return a valid kind Tenant struct",
			mock: func() {
				value, err := json.Marshal(tenant)
				if err != nil {
					t.Fatal(err)
					return
				}

				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewReader(value)),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, tenant, value.(*b3lbadmin.Tenant))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			config, err := API.GetTenant("localhost")
			test.validator(t, config, err)
		})
	}
}

func TestDeleteTenant(t *testing.T) {
	initTests()

	tests := []test{
		{
			name: "an error returned by rest client should return the error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("rest error")
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "http status != 204 should return an error",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       ioutil.NopCloser(bytes.NewReader([]byte("b3lb error"))),
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to delete tenant: b3lb error", err.Error())
			},
		},
		{
			name: "a valid request should return nil",
			mock: func() {
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNoContent,
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := API.DeleteTenant("localhost")
			test.validator(t, nil, err)
		})
	}
}

func TestApply(t *testing.T) {
	initTests()

	var kind string
	var resource interface{}

	tests := []test{
		{
			name: "an error returned by restclient while applying InstanceList should be returned",
			mock: func() {
				resource = &admin.InstanceList{
					Kind:      "InstanceList",
					Instances: map[string]string{},
				}
				kind = "InstanceList"
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("rest error")
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an error returned by restclient while applying Tenant should be returned",
			mock: func() {
				resource = &admin.Tenant{
					Kind:      "Tenant",
					Spec:      map[string]string{},
					Instances: []string{},
				}
				kind = "Tenant"
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("rest error")
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "an http status != 201 - Created should return an error",
			mock: func() {
				resource = &admin.Tenant{
					Kind:      "Tenant",
					Spec:      map[string]string{},
					Instances: []string{},
				}
				kind = "Tenant"
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusInternalServerError,
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "a valid request should return no error",
			mock: func() {
				resource = &admin.Tenant{
					Kind:      "Tenant",
					Spec:      map[string]string{},
					Instances: []string{},
				}
				kind = "Tenant"
				restclient.RestClientMockDoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusCreated,
					}, nil
				}
			},
			validator: func(t *testing.T, value interface{}, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := API.Apply(kind, &resource)
			test.validator(t, nil, err)
		})
	}
}
