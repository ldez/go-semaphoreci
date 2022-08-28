package v2

import (
	"net/http"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestTeamsService_GetByOrg(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/orgs/u/teams", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/TeamsService_GetByOrg.json")
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

	teams, _, err := client.Teams.GetByOrg("u")
	assert.NilError(t, err)

	expected := []Team{{
		ID:          "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:        "developers",
		Description: "Developers team",
		URL:         "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Permission:  "edit",
		UsersURL:    "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/users",
		ProjectsURL: "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		SecretsURL:  "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/secrets",
		CreatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, teams)
}

func TestTeamsService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/teams", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/TeamsService_GetByProject.json")
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

	teams, _, err := client.Teams.GetByProject("p")
	assert.NilError(t, err)

	expected := []Team{{
		ID:          "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:        "developers",
		Description: "Developers team",
		URL:         "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Permission:  "edit",
		UsersURL:    "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/users",
		ProjectsURL: "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		SecretsURL:  "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/secrets",
		CreatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, teams)
}

func TestTeamsService_GetBySecret(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/secrets/s/teams", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/TeamsService_GetBySecret.json")
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

	teams, _, err := client.Teams.GetBySecret("s")
	assert.NilError(t, err)

	expected := []Team{{
		ID:          "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:        "developers",
		Description: "Developers team",
		URL:         "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Permission:  "edit",
		UsersURL:    "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/users",
		ProjectsURL: "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		SecretsURL:  "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/secrets",
		CreatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, teams)
}

func TestTeamsService_Get(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/teams/t", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/TeamsService_Get.json")
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

	team, _, err := client.Teams.Get("t")
	assert.NilError(t, err)

	expected := &Team{
		ID:          "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:        "developers",
		Description: "Developers team",
		URL:         "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Permission:  "edit",
		UsersURL:    "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/users",
		ProjectsURL: "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		SecretsURL:  "https://api.semaphoreci.com/v2/teams/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/secrets",
		CreatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:   mustTimeParse("2017-06-10T16:59:51+02:00"),
	}

	assert.DeepEqual(t, expected, team)
}
