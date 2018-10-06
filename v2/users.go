package v2

import (
	"fmt"
	"net/http"
)

// UsersService http://semaphoreci.com/docs/api-v2-users.html
type UsersService service

// GetByOrg List all users for a organization
// http://semaphoreci.com/docs/api-v2-users.html#list-all-users-for-a-organization
func (c *UsersService) GetByOrg(orgUsername string) ([]User, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/users", orgUsername)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]User)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetByTeam List members of a team
// http://semaphoreci.com/docs/api-v2-users.html#list-members-of-a-team
func (c *UsersService) GetByTeam(teamID string) ([]User, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s/users", teamID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]User)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetByProject List all users for a project
// http://semaphoreci.com/docs/api-v2-users.html#list-all-users-for-a-project
func (c *UsersService) GetByProject(projectID string) ([]User, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/users", projectID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]User)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// TODO Add user to a team
// http://semaphoreci.com/docs/api-v2-users.html#add-user-to-a-team

// TODO Remove user from a team
// http://semaphoreci.com/docs/api-v2-users.html#remove-user-from-a-team
