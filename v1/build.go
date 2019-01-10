package v1

import (
	"fmt"
	"net/http"
	"net/url"
)

// BuildsService https://semaphoreci.com/docs/branches-and-builds-api.html
type BuildsService service

// GetInformation https://semaphoreci.com/docs/branches-and-builds-api.html#build_information
func (c *BuildsService) GetInformation(projectHashID string, branchID int, number int) (*BuildInformation, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v", projectHashID, branchID, number)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildInformation)

	_, err = c.client.Do(req, v)

	return v, err
}

// GetLog https://semaphoreci.com/docs/branches-and-builds-api.html#build_log
func (c *BuildsService) GetLog(projectHashID string, branchID int, number int) (*BuildLog, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v/log", projectHashID, branchID, number)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildLog)

	_, err = c.client.Do(req, v)

	return v, err
}

// RebuildLastRevision https://semaphoreci.com/docs/branches-and-builds-api.html#rebuild
func (c *BuildsService) RebuildLastRevision(projectHashID string, branchID int) (*BuildInformation, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/build", projectHashID, branchID)

	req, err := c.client.NewRequest(http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildInformation)

	_, err = c.client.Do(req, v)

	return v, err
}

// LaunchBuild https://semaphoreci.com/docs/branches-and-builds-api.html#launch_build
func (c *BuildsService) LaunchBuild(projectHashID string, branchID int, commitSHA string) (*BuildInformation, error) {
	u, err := url.Parse(fmt.Sprintf("projects/%s/%v/build?", projectHashID, branchID))
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("commit_sha", commitSHA)
	u.RawQuery = query.Encode()

	req, err := c.client.NewRequest(http.MethodPost, u.RawQuery, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildInformation)

	_, err = c.client.Do(req, v)

	return v, err
}

// Stop https://semaphoreci.com/docs/branches-and-builds-api.html#stop
func (c *BuildsService) Stop(projectHashID string, branchID int, number int) (*BuildInformation, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v/stop", projectHashID, branchID, number)

	req, err := c.client.NewRequest(http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BuildInformation)

	_, err = c.client.Do(req, v)

	return v, err
}

// DeployFromBuild https://semaphoreci.com/docs/branches-and-builds-api.html#deploy
func (c *BuildsService) DeployFromBuild(projectHashID string, branchID int, number int, serverID int) (*ServerStatus, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/builds/%v/deploy/%v", projectHashID, branchID, number, serverID)

	req, err := c.client.NewRequest(http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(ServerStatus)

	_, err = c.client.Do(req, v)

	return v, err
}
