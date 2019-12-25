package v2

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestConfigFilesService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/config_files", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ConfigFilesService_GetByProject.json")
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

	files, _, err := client.ConfigFiles.GetByProject("p")
	assert.NilError(t, err)

	expected := []ConfigFile{{
		ID:        "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Path:      ".credentials",
		URL:       "https://api.semaphoreci.com/v2/config_files/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Content:   "password: ec2c9f6f64b5",
		Shared:    true,
		CreatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, files)
}

func TestConfigFilesService_GetBySecret(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/secrets/s/config_files", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ConfigFilesService_GetBySecret.json")
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

	files, _, err := client.ConfigFiles.GetBySecret("s")
	assert.NilError(t, err)

	expected := []ConfigFile{{
		ID:        "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Path:      ".credentials",
		URL:       "https://api.semaphoreci.com/v2/config_files/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Content:   "password: ec2c9f6f64b5",
		Shared:    true,
		CreatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
	}}

	assert.DeepEqual(t, expected, files)
}

func TestConfigFilesService_Get(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/config_files/c", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ConfigFilesService_Get.json")
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

	file, _, err := client.ConfigFiles.Get("c")
	assert.NilError(t, err)

	expected := &ConfigFile{
		ID:        "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Path:      ".credentials",
		URL:       "https://api.semaphoreci.com/v2/config_files/86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Content:   "password: ec2c9f6f64b5",
		Shared:    true,
		CreatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
		UpdatedAt: mustTimeParse("2017-06-10T16:59:51+02:00"),
	}

	assert.DeepEqual(t, expected, file)
}
