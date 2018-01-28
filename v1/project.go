package v1

import (
	"net/http"
)

// https://semaphoreci.com/docs/projects-api.html

// GetProjects https://semaphoreci.com/docs/projects-api.html#
func (c Client) GetProjects() ([]Project, error) {
	req, err := c.newRequest(http.MethodGet, "projects", nil)
	if err != nil {
		return nil, err
	}

	v := new([]Project)

	_, err = do(req, v)

	return *v, err
}
