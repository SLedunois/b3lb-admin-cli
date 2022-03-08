package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lb/pkg/restclient"
	"github.com/SLedunois/b3lbctl/pkg/config"
)

// API is a public DefaultAdmin instance
var API Admin

const urlFormatter = "%s/admin/servers"

// Admin represents admin api interface
type Admin interface {
	List() ([]api.BigBlueButtonInstance, error)
	Add(url string, secret string) error
	Delete(instance string) error
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
	url := fmt.Sprintf(urlFormatter, *config.URL)
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

// Add performs a add admin call on b3lb
func (a *DefaultAdmin) Add(url string, secret string) error {
	apiURL := fmt.Sprintf(urlFormatter, *config.URL)
	instance := api.BigBlueButtonInstance{
		URL:    url,
		Secret: secret,
	}

	value, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	headers := authorization()
	headers["Content-Type"] = "application/json"
	resp, restErr := restclient.PostWithHeaders(apiURL, headers, value)
	if restErr == nil && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("api respond with a %d status code instead of %d", resp.StatusCode, http.StatusCreated)
	}

	return restErr
}

// Delete performs a delete admin call on B3LB
func (a *DefaultAdmin) Delete(instance string) error {
	apiURL := fmt.Sprintf(urlFormatter+"?url=%s", *config.URL, url.QueryEscape(instance))
	resp, restErr := restclient.DeleteWithHeaders(apiURL, authorization())
	if restErr == nil && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("api respond with a %d status code instead of %d", resp.StatusCode, http.StatusNoContent)
	}

	return restErr
}
