package v2

import (
	"fmt"
	"net/http"
)

// http://semaphoreci.com/docs/api-v2-config-files.html

// GetConfigFilesByProject List config files connected to a project
// http://semaphoreci.com/docs/api-v2-config-files.html#list-config-files-connected-to-a-project
func (c Client) GetConfigFilesByProject(projectID string) ([]ConfigFile, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/config_files", projectID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]ConfigFile)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetConfigFilesBySecret List config files belonging to a secret
// http://semaphoreci.com/docs/api-v2-config-files.html#list-config-files-belonging-to-a-secret
func (c Client) GetConfigFilesBySecret(secretID string) ([]ConfigFile, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s/config_files", secretID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]ConfigFile)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetConfigFile Get a config file
// http://semaphoreci.com/docs/api-v2-config-files.html#get-a-config-file
func (c Client) GetConfigFile(configFileID string) (*ConfigFile, *Response, error) {
	urlStr := fmt.Sprintf("config_files/%s", configFileID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(ConfigFile)

	resp, err := c.Do(req, v)

	return v, resp, err
}

// TODO Create a config file within a secret
// http://semaphoreci.com/docs/api-v2-config-files.html#create-a-config-file-within-a-secret

// TODO Update a config file
// http://semaphoreci.com/docs/api-v2-config-files.html#update-a-config-file

// TODO Delete a config file
// http://semaphoreci.com/docs/api-v2-config-files.html#delete-a-config-file
