package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SLedunois/b3lbctl/pkg/config"
	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lb/pkg/restclient"
)

// API is a public DefaultAdmin instance
var API Admin

// Admin represents admin api interface
type Admin interface {
	List() ([]api.BigBlueButtonInstance, error)
}

// DefaultAdmin is the default admin api struct. It an empty struct
type DefaultAdmin struct{}

// Init initialize a DefaultAdmin object
func Init() {
	API = &DefaultAdmin{}
}

func authorization() map[string]string {
	return map[string]string{
		"Authorization": *config.APIKey,
	}
}

// List performs a list admin call on b3lb
func (a *DefaultAdmin) List() ([]api.BigBlueButtonInstance, error) {
	url := fmt.Sprintf("%s/admin/servers", *config.URL)
	resp, err := restclient.GetWithHeaders(url, authorization())
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	instances := []api.BigBlueButtonInstance{}
	if err := json.Unmarshal(res, &instances); err != nil {
		return nil, err
	}

	return instances, nil
}
