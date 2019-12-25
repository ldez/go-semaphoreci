package v1

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestWebHooksService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/hooks", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/WebHooksService_GetByProject.json")
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

	webHooks, _, err := client.WebHooks.GetByProject("p")
	assert.NilError(t, err)

	expected := []WebHook{
		{
			ID:       256,
			URL:      "http://semaphoreflag.herokuapp.com/qH36J7zzMAxmF72f",
			HookType: "post_build",
		},
		{
			ID:       257,
			URL:      "http://semaphoreflag.herokuapp.com/qH36sdffgae2f",
			HookType: "post_deploy",
		},
	}

	assert.DeepEqual(t, expected, webHooks)
}

func TestWebHooksService_Create(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/hooks", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)

		content, err := ioutil.ReadFile("fixtures/WebHooksService_Create.json")
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

	webHook, _, err := client.WebHooks.Create("p", WebHook{
		URL:      "http://google.com",
		HookType: "all",
	})
	assert.NilError(t, err)

	expected := &WebHook{
		ID:       266,
		URL:      "http://google.com",
		HookType: "all",
	}

	assert.DeepEqual(t, expected, webHook)
}

func TestWebHooksService_Update(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/hooks/266", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPut)

		content, err := ioutil.ReadFile("fixtures/WebHooksService_Update.json")
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

	webHook, _, err := client.WebHooks.Update("p", 266, WebHook{
		URL:      "http://yahoo.com",
		HookType: "all",
	})
	assert.NilError(t, err)

	expected := &WebHook{
		ID:       266,
		URL:      "http://yahoo.com",
		HookType: "all",
	}

	assert.DeepEqual(t, expected, webHook)
}

func TestWebHooksService_Delete(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/hooks/266", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.WebHooks.Delete("p", 266)
	assert.NilError(t, err)

	assert.DeepEqual(t, resp.StatusCode, http.StatusNoContent)
}
