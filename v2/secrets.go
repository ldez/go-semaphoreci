package v2

import (
	"fmt"
	"net/http"
)

// http://semaphoreci.com/docs/api-v2-secrets.html

// GetSecretsByOrg List secrets in an organization
// http://semaphoreci.com/docs/api-v2-secrets.html#list-secrets-in-an-organization
func (c Client) GetSecretsByOrg(orgUsername string) ([]Secret, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s/secrets", orgUsername)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Secret)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetSecretsByTeam List secrets in a team
// http://semaphoreci.com/docs/api-v2-secrets.html#list-secrets-in-a-team
func (c Client) GetSecretsByTeam(teamID string) ([]Secret, *Response, error) {
	urlStr := fmt.Sprintf("teams/%s/secrets", teamID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Secret)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetSecretsByProject List secrets attached to a project
// http://semaphoreci.com/docs/api-v2-secrets.html#list-secrets-attached-to-a-project
func (c Client) GetSecretsByProject(projectID string) ([]Secret, *Response, error) {
	urlStr := fmt.Sprintf("projects/%s/secrets", projectID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Secret)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetSecret Get a secret
// http://semaphoreci.com/docs/api-v2-secrets.html#get-a-secret
func (c Client) GetSecret(secretID string) (*Secret, *Response, error) {
	urlStr := fmt.Sprintf("secrets/%s", secretID)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Secret)

	resp, err := c.Do(req, v)

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
