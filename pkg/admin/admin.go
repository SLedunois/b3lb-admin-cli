package admin

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/SLedunois/b3lb/pkg/api"
	"github.com/SLedunois/b3lb/pkg/balancer"
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
	ClusterStatus() ([]balancer.InstanceStatus, error)
	B3lbAPIStatus() (string, error)
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
	if restErr != nil {
		return restErr
	}

	if restErr == nil && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("api respond with a %d status code instead of %d", resp.StatusCode, http.StatusNoContent)
	}

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("instance does not found in your cluster")
	}

	return restErr
}

// ClusterStatus call cluster status admin api and return result
func (a *DefaultAdmin) ClusterStatus() ([]balancer.InstanceStatus, error) {
	resp, err := restclient.GetWithHeaders(fmt.Sprintf("%s/admin/api/cluster", *config.URL), authorization())
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	status := []balancer.InstanceStatus{}
	if err := json.Unmarshal(res, &status); err != nil {
		return nil, err
	}

	return status, nil
}

// B3lbAPIStatus returns the b3lb pi status
func (a *DefaultAdmin) B3lbAPIStatus() (string, error) {
	resp, err := restclient.Get(fmt.Sprintf("%s/bigbluebutton/api", *config.URL))
	if err != nil {
		return "", err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	status := &api.HealthCheck{}
	if err := xml.Unmarshal(res, &status); err != nil {
		return "", err
	}

	if status.ReturnCode == api.ReturnCodes().Success {
		return "Up", nil
	}

	return "Down", nil
}
