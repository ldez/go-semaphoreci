package v1

import (
	"net/http"
)

// Build status.
const (
	Passed  = "passed"
	Failed  = "failed"
	Stopped = "stopped"
	Pending = "pending"
)

// ProjectsService https://semaphoreci.com/docs/projects-api.html
type ProjectsService service

// Get https://semaphoreci.com/docs/projects-api.html#
func (c *ProjectsService) Get() ([]Project, *Response, error) {
	req, err := c.client.NewRequest(http.MethodGet, "projects", nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]Project)

	resp, err := c.client.Do(req, v)

	return *v, resp, err
}
