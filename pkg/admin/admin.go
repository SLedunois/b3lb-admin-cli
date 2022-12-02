package admin

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bigblueswarm/bbsctl/pkg/config"
	bbsadmin "github.com/bigblueswarm/bigblueswarm/v2/pkg/admin"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/api"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/balancer"
	bbsconfig "github.com/bigblueswarm/bigblueswarm/v2/pkg/config"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/restclient"
)

// API is a public DefaultAdmin instance
var API Admin

const urlFormatter = "%s/admin/api/instances"

// Admin represents admin api interface
type Admin interface {
	List() ([]api.BigBlueButtonInstance, error)
	Add(url string, secret string) error
	Delete(instance string) error
	ClusterStatus() ([]balancer.InstanceStatus, error)
	BBSAPIStatus() (string, error)
	GetConfiguration() (*bbsconfig.Config, error)
	GetTenants() (*bbsadmin.TenantList, error)
	GetTenant(hostname string) (*bbsadmin.Tenant, error)
	DeleteTenant(hostname string) error
	Apply(king string, resource *interface{}) error
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

// List performs a list admin call on bigblueswarm
func (a *DefaultAdmin) List() ([]api.BigBlueButtonInstance, error) {
	url := fmt.Sprintf(urlFormatter, *config.BBS)
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

// Add performs a add admin call on bigblueswarm
func (a *DefaultAdmin) Add(url string, secret string) error {
	apiURL := fmt.Sprintf(urlFormatter, *config.BBS)
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

// Delete performs a delete admin call on BigBlueSwarm
func (a *DefaultAdmin) Delete(instance string) error {
	apiURL := fmt.Sprintf(urlFormatter+"?url=%s", *config.BBS, url.QueryEscape(instance))
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
	resp, err := restclient.GetWithHeaders(fmt.Sprintf("%s/admin/api/cluster", *config.BBS), authorization())
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

// BBSAPIStatus returns the bigblueswarm pi status
func (a *DefaultAdmin) BBSAPIStatus() (string, error) {
	resp, err := restclient.Get(fmt.Sprintf("%s/bigbluebutton/api", *config.BBS))
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

// GetConfiguration return bigblueswarm configuration
func (a *DefaultAdmin) GetConfiguration() (*bbsconfig.Config, error) {
	resp, err := restclient.GetWithHeaders(fmt.Sprintf("%s/admin/api/configurations", *config.BBS), authorization())
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var config *bbsconfig.Config
	if err := json.Unmarshal(res, &config); err != nil {
		return nil, err
	}

	return config, nil
}

// GetTenants return bigblueswarm active tenants
func (a *DefaultAdmin) GetTenants() (*bbsadmin.TenantList, error) {
	resp, err := restclient.GetWithHeaders(fmt.Sprintf("%s/admin/api/tenants", *config.BBS), authorization())
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tenants *bbsadmin.TenantList
	if err := json.Unmarshal(res, &tenants); err != nil {
		return nil, err
	}

	return tenants, nil
}

// GetTenant return a specific tenant as kind Tenant
func (a *DefaultAdmin) GetTenant(hostname string) (*bbsadmin.Tenant, error) {
	resp, err := restclient.GetWithHeaders(fmt.Sprintf("%s/admin/api/tenants/%s", *config.BBS, hostname), authorization())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("tenant not found")
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("bigblueswarm internal error. Please check your bigblueswarm instance")
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tenant *bbsadmin.Tenant
	if err := json.Unmarshal(res, &tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

// DeleteTenant delete given tenant from bigblueswarm cluster
func (a *DefaultAdmin) DeleteTenant(hostname string) error {
	resp, err := restclient.DeleteWithHeaders(fmt.Sprintf("%s/admin/api/tenants/%s", *config.BBS, hostname), authorization())
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("bigblueswarm respond with an unreadable error: %s", err.Error())
		}

		return fmt.Errorf("unable to delete tenant: %s", string(res))
	}

	return nil
}

// Apply applies a resource to bigblueswarm cluster
func (a *DefaultAdmin) Apply(kind string, resource *interface{}) error {
	var url string

	if kind == "InstanceList" {
		url = fmt.Sprintf("%s/admin/api/instances", *config.BBS)
	} else {
		url = fmt.Sprintf("%s/admin/api/tenants", *config.BBS)
	}

	b, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	resp, err := restclient.PostWithHeaders(url, authorization(), b)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("bigblueswarm returns a %d status instead of %d", resp.StatusCode, http.StatusCreated)
	}

	return nil
}
