package v1

import (
	"fmt"
	"net/http"
)

// ServersService https://semaphoreci.com/docs/servers-and-deploys-api.html
type ServersService service

// GetByProject https://semaphoreci.com/docs/servers-and-deploys-api.html#project_servers
func (c *ServersService) GetByProject(projectHashID string) ([]Server, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/server", projectHashID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Server)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetStatus https://semaphoreci.com/docs/servers-and-deploys-api.html#server_status
func (c *ServersService) GetStatus(projectHashID string, serverID int) (*ServerStatus, error) {
	urlStr := fmt.Sprintf("projects/%s/servers/%v/status", projectHashID, serverID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(ServerStatus)

	_, err = c.client.Do(req, v)

	return v, err
}

// GetHistory https://semaphoreci.com/docs/servers-and-deploys-api.html#server_history
func (c *ServersService) GetHistory(projectHashID string, serverID int, opts *ServerHistoryOptions) (*ServerDeploy, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/servers/%v", projectHashID, serverID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(ServerDeploy)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}
