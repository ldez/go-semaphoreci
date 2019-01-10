package v2

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestOrganizationsService_GetYours(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/orgs", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/OrganizationsService_GetYours.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	organizations, _, err := client.Organizations.GetYours()
	assert.NilError(t, err)

	expected := []Organization{{
		ID:          "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:        "Rendered Text",
		URL:         "https://api.semaphoreci.com/v2/orgs/renderedtext",
		ProjectsURL: "https://api.semaphoreci.com/v2/orgs/renderedtext/projects",
		SecretsURL:  "https://api.semaphoreci.com/v2/orgs/renderedtext/secrets",
		UsersURL:    "https://api.semaphoreci.com/v2/orgs/renderedtext/users",
		TeamsURL:    "https://api.semaphoreci.com/v2/orgs/renderedtext/teams",
		Username:    "renderedtext",
		CreatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, organizations)
}

func TestOrganizationsService_Get(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/orgs/u", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/OrganizationsService_Get.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	organization, _, err := client.Organizations.Get("u")
	assert.NilError(t, err)

	expected := &Organization{
		ID:          "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:        "Rendered Text",
		URL:         "https://api.semaphoreci.com/v2/orgs/renderedtext",
		ProjectsURL: "https://api.semaphoreci.com/v2/orgs/renderedtext/projects",
		SecretsURL:  "https://api.semaphoreci.com/v2/orgs/renderedtext/secrets",
		UsersURL:    "https://api.semaphoreci.com/v2/orgs/renderedtext/users",
		TeamsURL:    "https://api.semaphoreci.com/v2/orgs/renderedtext/teams",
		Username:    "renderedtext",
		CreatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
	}

	assert.DeepEqual(t, expected, organization)
}
