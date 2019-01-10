package v1

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestServersService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/server", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ServersService_GetByProject.json")
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

	servers, _, err := client.Servers.GetByProject("p")
	assert.NilError(t, err)

	expected := []Server{
		{
			ID:        9,
			Name:      "staging",
			ServerURL: "https://semaphoreci.com/api/v1/projects/:hash_id/servers/9/status?auth_token=:auth_token",
		},
		{
			ID:        11,
			Name:      "production",
			ServerURL: "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/status?auth_token=:auth_token",
		},
	}
	assert.DeepEqual(t, expected, servers)
}

func TestServersService_GetStatus(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/servers/666/status", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ServersService_GetStatus.json")
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

	statuses, err := client.Servers.GetStatus("p", 666)
	assert.NilError(t, err)

	expected := &ServerStatus{
		ServerName:       "production",
		ServerURL:        "https://semaphoreci.com/projects/1/servers/11",
		ServerStatusURL:  "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/status?auth_token=:auth_token",
		ServerHistoryURL: "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11?auth_token=:auth_token",
		DeploymentMethod: "capistrano",
		Strategy:         "manual",
		BranchName:       "",
		ProjectName:      "semaphore",
		Number:           27,
		Result:           "passed",
		CreatedAt:        mustTimeParse("2013-07-23T10:57:42+02:00"),
		UpdatedAt:        mustTimeParse("2013-07-23T11:00:41+02:00"),
		StartedAt:        mustTimeParse("2013-07-23T10:57:49+02:00"),
		FinishedAt:       mustTimeParse("2013-07-23T11:00:41+02:00"),
		HTMLURL:          "https://semaphoreci.com/projects/1/servers/11/deploys/27",
		DeployURL:        "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/deploys/27?auth_token=:auth_token",
		DeployLogURL:     "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/deploys/27/log?auth_token=:auth_token",
		BuildURL:         "https://semaphoreci.com/api/v1/projects/:hash_id/29803/builds/119?auth_token=:auth_token",
		BuildHTMLURL:     "https://semaphoreci.com/projects/1/branches/29803/builds/119",
		Commit: &Commit{
			ID:          "222f031528545f2bc8284a4725fe160a0fb443x1",
			URL:         "https://github.com/renderedtext/semaphore/commit/222f031528545f2bc8284a4725fe160a0fb443x1",
			AuthorName:  "Marko Anastasov",
			AuthorEmail: "marko@renderedtext.com",
			Message:     "Merge branch 'staging'",
			Timestamp:   mustTimeParse("2013-07-22T17:52:27+02:00"),
		},
	}

	assert.DeepEqual(t, expected, statuses)
}

func TestServersService_GetHistory(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/servers/666", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ServersService_GetHistory.json")
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

	history, _, err := client.Servers.GetHistory("p", 666, &ServerHistoryOptions{Page: 1})
	assert.NilError(t, err)

	expected := &ServerDeploy{
		ServerName:       "production",
		ServerURL:        "https://semaphoreci.com/projects/1/servers/11",
		ServerStatusURL:  "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/status?auth_token=:auth_token",
		ServerHistoryURL: "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11?auth_token=:auth_token",
		DeploymentMethod: "capistrano",
		Strategy:         "manual",
		BranchName:       "",
		ProjectName:      "semaphore",
		Deploys: []Deploy{
			{
				ProjectName:  "semaphore",
				ServerName:   "production",
				Number:       27,
				Result:       "passed",
				CreatedAt:    mustTimeParse("2013-07-23T10:57:42+02:00"),
				UpdatedAt:    mustTimeParse("2013-07-23T11:00:41+02:00"),
				StartedAt:    mustTimeParse("2013-07-23T10:57:49+02:00"),
				FinishedAt:   mustTimeParse("2013-07-23T11:00:41+02:00"),
				HTMLURL:      "https://semaphoreci.com/projects/1/servers/11/deploys/27",
				DeployURL:    "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/deploys/27?auth_token=:auth_token",
				DeployLogURL: "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/deploys/27/log?auth_token=:auth_token",
				BuildURL:     "https://semaphoreci.com/api/v1/projects/:hash_id/29803/builds/119?auth_token=:auth_token",
				BuildHTMLURL: "https://semaphoreci.com/projects/1/branches/29803/builds/119",
				Commit: &Commit{
					ID:          "222f0123418545f21234184a4725fe16asfa125123",
					URL:         "https://github.com/renderedtext/semaphore/commit/222f0123418545f21234184a4725fe16asfa125123",
					AuthorName:  "Marko Anastasov",
					AuthorEmail: "marko@renderedtext.com",
					Message:     "Merge branch 'staging'",
					Timestamp:   mustTimeParse("2013-07-22T17:52:27+02:00"),
				},
			},
			{
				ProjectName:  "semaphore",
				ServerName:   "production",
				Number:       26,
				Result:       "passed",
				CreatedAt:    mustTimeParse("2013-07-23T09:57:42+02:00"),
				UpdatedAt:    mustTimeParse("2013-07-23T10:00:41+02:00"),
				StartedAt:    mustTimeParse("2013-07-23T09:57:49+02:00"),
				FinishedAt:   mustTimeParse("2013-07-23T10:00:41+02:00"),
				HTMLURL:      "https://semaphoreci.com/projects/1/servers/11/deploys/27",
				DeployURL:    "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/deploys/27?auth_token=:auth_token",
				DeployLogURL: "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/deploys/27/log?auth_token=:auth_token",
				BuildURL:     "https://semaphoreci.com/api/v1/projects/:hash_id/29803/builds/119?auth_token=:auth_token",
				BuildHTMLURL: "https://semaphoreci.com/projects/1/branches/29803/builds/119",
				Commit: &Commit{
					ID:          "222f0123418545f21234184a4725fe16asfa125123",
					URL:         "https://github.com/renderedtext/semaphore/commit/222f0123418545f21234184a4725fe16asfa125123",
					AuthorName:  "Marko Anastasov",
					AuthorEmail: "marko@renderedtext.com",
					Message:     "Merge branch 'staging'",
					Timestamp:   mustTimeParse("2013-07-22T16:52:27+02:00"),
				},
			},
		},
	}

	assert.DeepEqual(t, expected, history)
}
