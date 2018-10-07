package v1

import (
	"fmt"
	"net/http"
)

// BranchesService https://semaphoreci.com/docs/branches-and-builds-api.html
type BranchesService service

// GetByProject https://semaphoreci.com/docs/branches-and-builds-api.html#project_branches
func (c *BranchesService) GetByProject(projectHashID string) ([]Branch, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/branches", projectHashID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Branch)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetStatus https://semaphoreci.com/docs/branches-and-builds-api.html#branch_status
func (c *BranchesService) GetStatus(projectHashID string, branchID int) (*BranchStatus, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/status", projectHashID, branchID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BranchStatus)

	_, err = c.client.Do(req, v)

	return v, err
}

// GetHistory https://semaphoreci.com/docs/branches-and-builds-api.html#branch_history
func (c *BranchesService) GetHistory(projectHashID string, branchID int, opts *BranchHistoryOptions) (*BranchHistory, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/%v", projectHashID, branchID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(BranchHistory)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}
