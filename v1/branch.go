package v1

import (
	"fmt"
	"net/http"
)

// https://semaphoreci.com/docs/branches-and-builds-api.html

// GetBranches https://semaphoreci.com/docs/branches-and-builds-api.html#project_branches
func (c Client) GetBranches(projectHashID string) ([]Branch, error) {
	urlStr := fmt.Sprintf("projects/%s/branches", projectHashID)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new([]Branch)

	_, err = do(req, v)

	return *v, err
}

// GetBranchStatus https://semaphoreci.com/docs/branches-and-builds-api.html#branch_status
func (c Client) GetBranchStatus(projectHashID string, branchID int) (*BranchStatus, error) {
	urlStr := fmt.Sprintf("projects/%s/%v/status", projectHashID, branchID)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(BranchStatus)

	_, err = do(req, v)

	return v, err
}

// GetBranchHistory https://semaphoreci.com/docs/branches-and-builds-api.html#branch_history
func (c Client) GetBranchHistory(projectHashID string, branchID int, opts *BranchHistoryOptions) (*BranchHistory, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/%v", projectHashID, branchID)

	req, err := c.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(BranchHistory)

	resp, err := do(req, v)

	return v, resp, err
}
