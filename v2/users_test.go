package v2

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestUsersService_GetByOrg(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/orgs/u/users", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/UsersService_GetByOrg.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	users, _, err := client.Users.GetByOrg("u")
	assert.NilError(t, err)

	expected := []User{{
		UID:       "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Username:  "johndoe",
		Name:      "John Doe",
		CreatedAt: mustTimeParse("2017-06-11T16:59:50+02:00"),
		UpdatedAt: mustTimeParse("2017-06-11T16:59:50+02:00"),
	}}

	assert.DeepEqual(t, expected, users)
}

func TestUsersService_GetByTeam(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/teams/t/users", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/UsersService_GetByTeam.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	users, _, err := client.Users.GetByTeam("t")
	assert.NilError(t, err)

	expected := []User{{
		UID:       "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Username:  "johndoe",
		Name:      "John Doe",
		CreatedAt: mustTimeParse("2017-06-11T16:59:50+02:00"),
		UpdatedAt: mustTimeParse("2017-06-11T16:59:50+02:00"),
	}}

	assert.DeepEqual(t, expected, users)
}

func TestUsersService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/users", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/UsersService_GetByProject.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	users, _, err := client.Users.GetByProject("p")
	assert.NilError(t, err)

	expected := []User{{
		UID:       "86e78b7e-2f9c-45a7-9939-ec2c9f6f64b5",
		Username:  "johndoe",
		Name:      "John Doe",
		CreatedAt: mustTimeParse("2017-06-11T16:59:50+02:00"),
		UpdatedAt: mustTimeParse("2017-06-11T16:59:50+02:00"),
	}}

	assert.DeepEqual(t, expected, users)
}
