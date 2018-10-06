package v2

import (
	"fmt"
	"net/http"
)

// http://semaphoreci.com/docs/api-v2-users.html

// GetUsersByOrg List all users for a organization
// http://semaphoreci.com/docs/api-v2-users.html#list-all-users-for-a-organization
func (c Client) GetUsersByOrg(orgUsername string) ([]User, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/users", orgUsername)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]User)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetUsersByTeam List members of a team
// http://semaphoreci.com/docs/api-v2-users.html#list-members-of-a-team
func (c Client) GetUsersByTeam(teamID string) ([]User, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s/users", teamID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]User)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetUsersByProject List all users for a project
// http://semaphoreci.com/docs/api-v2-users.html#list-all-users-for-a-project
func (c Client) GetUsersByProject(projectID string) ([]User, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/users", projectID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]User)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// TODO Add user to a team
// http://semaphoreci.com/docs/api-v2-users.html#add-user-to-a-team

// TODO Remove user from a team
// http://semaphoreci.com/docs/api-v2-users.html#remove-user-from-a-team
