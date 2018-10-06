package v2

import (
	"fmt"
	"net/http"
)

// ProjectsService http://semaphoreci.com/docs/api-v2-projects.html
type ProjectsService service

// GetByOrg List projects in an organization
// http://semaphoreci.com/docs/api-v2-projects.html#list-projects-in-an-organization
func (c *ProjectsService) GetByOrg(orgUsername string) ([]Project, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/projects", orgUsername)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Project)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetByTeam List project added to a team
// http://semaphoreci.com/docs/api-v2-projects.html#list-project-added-to-a-team
func (c *ProjectsService) GetByTeam(teamID string) ([]Project, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s/projects", teamID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Project)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetBySecret List projects for a secret
// http://semaphoreci.com/docs/api-v2-projects.html#list-projects-for-a-secret
func (c *ProjectsService) GetBySecret(secretID string) ([]Project, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s/projects", secretID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Project)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// TODO Create a project in an organization
// http://semaphoreci.com/docs/api-v2-projects.html#create-a-project-in-an-organization

// TODO Add project to a team
// http://semaphoreci.com/docs/api-v2-projects.html#add-project-to-a-team

// TODO Remove project from a team
// http://semaphoreci.com/docs/api-v2-projects.html#remove-project-from-a-team
