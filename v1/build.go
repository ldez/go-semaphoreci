package v1

import (
	"fmt"
	"net/http"
)

// https://semaphoreci.com/docs/branches-and-builds-api.html

// GetBuildInformation https://semaphoreci.com/docs/branches-and-builds-api.html#build_information
func (c Client) GetBuildInformation(projectHashID string, branchID int, number int) (*BuildInformation, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v", projectHashID, branchID, number)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildInformation)

	_, err = do(req, v)

	return v, err
}

// GetBuildLog https://semaphoreci.com/docs/branches-and-builds-api.html#build_log
func (c Client) GetBuildLog(projectHashID string, branchID int, number int) (*BuildLog, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v/log", projectHashID, branchID, number)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildLog)

	_, err = do(req, v)

	return v, err
}

// RebuildLastRevision https://semaphoreci.com/docs/branches-and-builds-api.html#rebuild
func (c Client) RebuildLastRevision(projectHashID string, branchID int) (*BuildInformation, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/build", projectHashID, branchID)

	req, err := c.newRequest(http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildInformation)

	_, err = do(req, v)

	return v, err
}

// StopBuild https://semaphoreci.com/docs/branches-and-builds-api.html#stop
func (c Client) StopBuild(projectHashID string, branchID int, number int) (*BuildInformation, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v/stop", projectHashID, branchID, number)

	req, err := c.newRequest(http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildInformation)

	_, err = do(req, v)

	return v, err
}

// DeployFromBuild https://semaphoreci.com/docs/branches-and-builds-api.html#deploy
func (c Client) DeployFromBuild(projectHashID string, branchID int, number int, serverID int) (*ServerStatus, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v/deploy/%v", projectHashID, branchID, number, serverID)

	req, err := c.newRequest(http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(ServerStatus)

	_, err = do(req, v)

	return v, err
}
