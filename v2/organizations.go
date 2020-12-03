package v2

import (
	"fmt"
	"net/http"
)

// OrganizationsService http://semaphoreci.com/docs/api-v2-orgs.html
type OrganizationsService service

// GetYours List your organizations.
// http://semaphoreci.com/docs/api-v2-orgs.html#list-your-organizations
func (c *OrganizationsService) GetYours() ([]Organization, *Response, error) {
	req, err := c.client.NewRequest(http.MethodGet, "orgs", nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Organization)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}

// Get Get an organization.
// http://semaphoreci.com/docs/api-v2-orgs.html#get-an-organization
func (c *OrganizationsService) Get(username string) (*Organization, *Response, error) {
	urlStr := fmt.Sprintf("orgs/%s", username)

	req, err := c.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(Organization)

	resp, err := c.client.Do(req, v)

	return v, resp, err
}
