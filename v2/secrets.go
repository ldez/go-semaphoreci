package v2

import (
	"fmt"
	"net/http"
)

// SecretsService http://semaphoreci.com/docs/api-v2-secrets.html
type SecretsService service

// GetByOrg List secrets in an organization.
// http://semaphoreci.com/docs/api-v2-secrets.html#list-secrets-in-an-organization
func (c *SecretsService) GetByOrg(orgUsername string) ([]Secret, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/secrets", orgUsername)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Secret)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetByTeam List secrets in a team.
// http://semaphoreci.com/docs/api-v2-secrets.html#list-secrets-in-a-team
func (c *SecretsService) GetByTeam(teamID string) ([]Secret, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s/secrets", teamID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Secret)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// GetByProject List secrets attached to a project.
// http://semaphoreci.com/docs/api-v2-secrets.html#list-secrets-attached-to-a-project
func (c *SecretsService) GetByProject(projectID string) ([]Secret, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/secrets", projectID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Secret)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// Get Get a secret.
// http://semaphoreci.com/docs/api-v2-secrets.html#get-a-secret
func (c *SecretsService) Get(secretID string) (*Secret, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s", secretID)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Secret)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}

// TODO Create secret in an organization
// http://semaphoreci.com/docs/api-v2-secrets.html#create-secret-in-an-organization

// TODO Update a secret
// http://semaphoreci.com/docs/api-v2-secrets.html#update-a-secret

// TODO Delete a secret
// http://semaphoreci.com/docs/api-v2-secrets.html#delete-a-secret

// TODO Add a secret to a team
// http://semaphoreci.com/docs/api-v2-secrets.html#add-a-secret-to-a-team

// TODO Attach a secret to a project
// http://semaphoreci.com/docs/api-v2-secrets.html#attach-a-secret-to-a-project

// TODO Remove secret from a team
// http://semaphoreci.com/docs/api-v2-secrets.html#remove-secret-from-a-team

// TODO Detach a secret from a project
// http://semaphoreci.com/docs/api-v2-secrets.html#dettach-a-secret-from-a-project
