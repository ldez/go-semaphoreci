package v1

import (
	"fmt"
	"net/http"
)

// GetServers https://semaphoreci.com/docs/servers-and-deploys-api.html#project_servers
func (c Client) GetServers(projectHashID string) ([]Server, error) {
	urlStr := fmt.Sprintf("projects/%s/server", projectHashID)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new([]Server)

	_, err = do(req, v)

	return *v, err
}

// GetServerStatus https://semaphoreci.com/docs/servers-and-deploys-api.html#server_status
func (c Client) GetServerStatus(projectHashID string, serverID int) (*ServerStatus, error) {
	urlStr := fmt.Sprintf("projects/%s/servers/%v/status", projectHashID, serverID)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(ServerStatus)

	_, err = do(req, v)

	return v, err
}

// GetServerHistory https://semaphoreci.com/docs/servers-and-deploys-api.html#server_history
func (c Client) GetServerHistory(projectHashID string, serverID int) (*ServerDeploy, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/servers/%v", projectHashID, serverID)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(ServerDeploy)

	resp, err := do(req, v)

	return v, resp, err
}
