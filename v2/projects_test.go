package v2

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestProjectsService_GetByOrg(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/orgs/u/projects", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ProjectsService_GetByOrg.json")
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

	projects, _, err := client.Projects.GetByOrg("u")
	assert.NilError(t, err)

	expected := []Project{{
		ID:         "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:       "elixir-lang",
		HTMLURL:    "https://semaphoreci.com/renderedtext/elixir-lang",
		UsersURL:   "https://api.semaphoreci.com/v2/projects/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/users",
		SecretsURL: "https://api.semaphoreci.com/v2/projects/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/secrets",
		UpdatedAt:  mustTimeParse("2017-06-10T16:59:51+02:00"),
		CreatedAt:  mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, projects)
}

func TestProjectsService_GetByTeam(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/teams/t/projects", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ProjectsService_GetByTeam.json")
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

	projects, _, err := client.Projects.GetByTeam("t")
	assert.NilError(t, err)

	expected := []Project{{
		ID:         "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:       "elixir-lang",
		HTMLURL:    "https://semaphoreci.com/renderedtext/elixir-lang",
		UsersURL:   "https://api.semaphoreci.com/v2/projects/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/users",
		SecretsURL: "https://api.semaphoreci.com/v2/projects/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/secrets",
		UpdatedAt:  mustTimeParse("2017-06-10T16:59:51+02:00"),
		CreatedAt:  mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, projects)
}

func TestProjectsService_GetBySecret(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/secrets/s/projects", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ProjectsService_GetBySecret.json")
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

	projects, _, err := client.Projects.GetBySecret("s")
	assert.NilError(t, err)

	expected := []Project{{
		ID:         "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:       "elixir-lang",
		HTMLURL:    "https://semaphoreci.com/renderedtext/elixir-lang",
		UsersURL:   "https://api.semaphoreci.com/v2/projects/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/users",
		SecretsURL: "https://api.semaphoreci.com/v2/projects/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/secrets",
		UpdatedAt:  mustTimeParse("2017-06-10T16:59:51+02:00"),
		CreatedAt:  mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, projects)
}
