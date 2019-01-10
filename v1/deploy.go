package v1

import (
	"fmt"
	"net/http"
)

// DeploysService https://semaphoreci.com/docs/servers-and-deploys-api.html
type DeploysService service

// GetInformation https://semaphoreci.com/docs/servers-and-deploys-api.html#deploy_information
func (c *DeploysService) GetInformation(projectHashID string, serverID int, number int) (*Deploy, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/servers/%v/deploys/%v", projectHashID, serverID, number)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Deploy)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}

// GetLog https://semaphoreci.com/docs/servers-and-deploys-api.html#deploy_log
func (c *DeploysService) GetLog(projectHashID string, serverID int, number int) (*DeployLog, error) {
	// GET /api/v1/projects/:hash_id/servers/:id/deploys/:number/log

	urlStr := fmt.Sprintf("projects/%s/servers/%v/deploys/%v/log", projectHashID, serverID, number)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(DeployLog)

	_, err = c.client.Do(req, v)

	return v, err
}

// Stop https://semaphoreci.com/docs/servers-and-deploys-api.html#stop
func (c *DeploysService) Stop(projectHashID string, serverID int, number int) (*Deploy, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/servers/%v/deploys/%v/stop", projectHashID, serverID, number)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Deploy)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}
