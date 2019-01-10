package v2

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestEnvVarsService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/env_vars", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/EnvVarsService_GetByProject.json")
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

	envVars, _, err := client.EnvVars.GetByProject("p")
	assert.NilError(t, err)

	expected := []EnvVar{{
		ID:        "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:      "API_TOKEN",
		URL:       "https://api.semaphoreci.com/v2/env_vars/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Content:   "8CgLAxXn",
		Shared:    true,
		CreatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, envVars)
}

func TestEnvVarsService_GetBySecret(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/secrets/s/env_vars", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/EnvVarsService_GetBySecret.json")
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

	envVars, _, err := client.EnvVars.GetBySecret("s")
	assert.NilError(t, err)

	expected := []EnvVar{{
		ID:        "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:      "API_TOKEN",
		URL:       "https://api.semaphoreci.com/v2/env_vars/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Content:   "8CgLAxXn",
		Shared:    true,
		CreatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, envVars)
}

func TestEnvVarsService_Get(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/env_vars/e", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/EnvVarsService_Get.json")
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

	envVar, _, err := client.EnvVars.Get("e")
	assert.NilError(t, err)

	expected := &EnvVar{
		ID:        "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Name:      "API_TOKEN",
		URL:       "https://api.semaphoreci.com/v2/env_vars/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Content:   "8CgLAxXn",
		Shared:    true,
		CreatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
	}

	assert.DeepEqual(t, expected, envVar)
}
