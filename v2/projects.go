package v2

import (
	"fmt"
	"net/http"
)

// http://semaphoreci.com/docs/api-v2-projects.html

// GetProjectsByOrg List projects in an organization
// http://semaphoreci.com/docs/api-v2-projects.html#list-projects-in-an-organization
func (c Client) GetProjectsByOrg(orgUsername string) ([]Project, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/projects", orgUsername)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Project)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetProjectsByTeam List project added to a team
// http://semaphoreci.com/docs/api-v2-projects.html#list-project-added-to-a-team
func (c Client) GetProjectsByTeam(teamID string) ([]Project, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s/projects", teamID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Project)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetProjectsBySecret List projects for a secret
// http://semaphoreci.com/docs/api-v2-projects.html#list-projects-for-a-secret
func (c Client) GetProjectsBySecret(secretID string) ([]Project, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s/projects", secretID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Project)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// TODO Create a project in an organization
// http://semaphoreci.com/docs/api-v2-projects.html#create-a-project-in-an-organization

// TODO Add project to a team
// http://semaphoreci.com/docs/api-v2-projects.html#add-project-to-a-team

// TODO Remove project from a team
// http://semaphoreci.com/docs/api-v2-projects.html#remove-project-from-a-team
