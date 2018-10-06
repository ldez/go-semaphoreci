package v2

import (
	"fmt"
	"net/http"
)

// TeamsService http://semaphoreci.com/docs/api-v2-teams.html
type TeamsService service

// GetByOrg List teams in an organization
// http://semaphoreci.com/docs/api-v2-teams.html#list-teams-in-an-organization
func (c *TeamsService) GetByOrg(orgUsername string) ([]Team, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/teams", orgUsername)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Team)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetByProject List all teams connected to project
// http://semaphoreci.com/docs/api-v2-teams.html#list-all-teams-connected-to-project
func (c *TeamsService) GetByProject(projectID string) ([]Team, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/teams", projectID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Team)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetBySecret List teams for a secret
// http://semaphoreci.com/docs/api-v2-teams.html#list-teams-for-a-secret
func (c *TeamsService) GetBySecret(secretID string) ([]Team, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s/teams", secretID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Team)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// Get Get a team
// http://semaphoreci.com/docs/api-v2-teams.html#get-a-team
func (c *TeamsService) Get(teamID string) (*Team, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s", teamID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Team)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}

// TODO Create a team in an organization
// http://semaphoreci.com/docs/api-v2-teams.html#create-a-team-in-an-organization

// TODO Update a team
// http://semaphoreci.com/docs/api-v2-teams.html#update-a-team

// TODO Delete a team
// http://semaphoreci.com/docs/api-v2-teams.html#delete-a-team
