package v2

import (
	"fmt"
	"net/http"
)

// http://semaphoreci.com/docs/api-v2-teams.html

// GetTeamsByOrg List teams in an organization
// http://semaphoreci.com/docs/api-v2-teams.html#list-teams-in-an-organization
func (c Client) GetTeamsByOrg(orgUsername string) ([]Team, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/teams", orgUsername)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Team)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetTeamsByProject List all teams connected to project
// http://semaphoreci.com/docs/api-v2-teams.html#list-all-teams-connected-to-project
func (c Client) GetTeamsByProject(projectID string) ([]Team, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/teams", projectID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Team)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetTeamsBySecret List teams for a secret
// http://semaphoreci.com/docs/api-v2-teams.html#list-teams-for-a-secret
func (c Client) GetTeamsBySecret(secretID string) ([]Team, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s/teams", secretID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Team)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetTeam Get a team
// http://semaphoreci.com/docs/api-v2-teams.html#get-a-team
func (c Client) GetTeam(teamID string) (*Team, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s", teamID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Team)

	resp, err := c.Do(req, v)

	return v, resp, err
}

// TODO Create a team in an organization
// http://semaphoreci.com/docs/api-v2-teams.html#create-a-team-in-an-organization

// TODO Update a team
// http://semaphoreci.com/docs/api-v2-teams.html#update-a-team

// TODO Delete a team
// http://semaphoreci.com/docs/api-v2-teams.html#delete-a-team
