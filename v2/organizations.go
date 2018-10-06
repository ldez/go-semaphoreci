package v2

import (
	"fmt"
	"net/http"
)

// http://semaphoreci.com/docs/api-v2-orgs.html

// GetOrganizations List your organizations
// http://semaphoreci.com/docs/api-v2-orgs.html#list-your-organizations
func (c Client) GetOrganizations(orgUsername string) ([]Organization, *Response, error) {
	req, err := c.NewRequest(http.MethodGet, "orgs", nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Organization)

	resp, err := c.Do(req, v)

	return *v, resp, err
}

// GetOrganization Get an organization
// http://semaphoreci.com/docs/api-v2-orgs.html#get-an-organization
func (c Client) GetOrganization(username string) (*Organization, error) {
	urlStr := fmt.Sprintf("orgs/%s/projects", username)

	req, err := c.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	v := new(Organization)

	_, err = c.Do(req, v)

	return v, err
}
