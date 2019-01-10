package v2

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestSecretsService_GetByOrg(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/orgs/u/secrets", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/SecretsService_GetByOrg.json")
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

	secrets, _, err := client.Secrets.GetByOrg("u")
	assert.NilError(t, err)

	expected := []Secret{{
		ID:             "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:           "AWS Tokens",
		Description:    "AWS tokens for deployment",
		URL:            "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		ProjectsURL:    "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		TeamsURL:       "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/teams",
		EnvVarsURL:     "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/env_vars",
		ConfigFilesURL: "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/config_files",
		CreatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, secrets)
}

func TestSecretsService_GetByTeam(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/teams/t/secrets", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/SecretsService_GetByTeam.json")
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

	secrets, _, err := client.Secrets.GetByTeam("t")
	assert.NilError(t, err)

	expected := []Secret{{
		ID:             "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:           "AWS Tokens",
		Description:    "AWS tokens for deployment",
		URL:            "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		ProjectsURL:    "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		TeamsURL:       "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/teams",
		EnvVarsURL:     "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/env_vars",
		ConfigFilesURL: "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/config_files",
		CreatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, secrets)
}

func TestSecretsService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/secrets", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/SecretsService_GetByProject.json")
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

	secrets, _, err := client.Secrets.GetByProject("p")
	assert.NilError(t, err)

	expected := []Secret{{
		ID:             "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:           "AWS Tokens",
		Description:    "AWS tokens for deployment",
		URL:            "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		ProjectsURL:    "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		TeamsURL:       "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/teams",
		EnvVarsURL:     "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/env_vars",
		ConfigFilesURL: "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/config_files",
		CreatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, secrets)
}

func TestSecretsService_Get(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/secrets/s", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/SecretsService_Get.json")
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

	secret, _, err := client.Secrets.Get("s")
	assert.NilError(t, err)

	expected := &Secret{
		ID:             "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:           "AWS Tokens",
		Description:    "AWS tokens for deployment",
		URL:            "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		ProjectsURL:    "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/projects",
		TeamsURL:       "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/teams",
		EnvVarsURL:     "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/env_vars",
		ConfigFilesURL: "https://api.semaphoreci.com/v2/secrets/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5/config_files",
		CreatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt:      mustTimeParse("2017-06-10T16:59:51+02:00"),
	}

	assert.DeepEqual(t, expected, secret)
}
