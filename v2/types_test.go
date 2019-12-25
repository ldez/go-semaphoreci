package v2

import (
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestResponse_populatePageValues(t *testing.T) {
	response := Response{
		Response: &http.Response{
			Header: map[string][]string{
				"Total":    {"4"},
				"Per-Page": {"30"},
				"Link":     {`<http://api.semaphoreci.com/v2/orgs?page=1>; rel="first", <http://api.semaphoreci.com/v2/orgs?page=2>; rel="last"`},
			},
		},
	}

	response.populatePageValues()

	assert.Equal(t, 4, response.Total)
	assert.Equal(t, 30, response.PerPage)
	assert.Equal(t, 0, response.NextPage)
	assert.Equal(t, 0, response.PrevPage)
	assert.Equal(t, 1, response.FirstPage)
	assert.Equal(t, 2, response.LastPage)
}
