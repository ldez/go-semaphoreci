package v1

import (
	"net/http"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDeploysService_GetInformation(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/servers/666/deploys/123", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/DeploysService_GetInformation.json")
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

	info, _, err := client.Deploys.GetInformation("p", 666, 123)
	assert.NilError(t, err)

	expected := &Deploy{
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
	}

	assert.DeepEqual(t, expected, info)
}

func TestDeploysService_GetLog(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/servers/666/deploys/123/log", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/DeploysService_GetLog.json")
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

	info, err := client.Deploys.GetLog("p", 666, 123)
	assert.NilError(t, err)

	expected := &DeployLog{
		Threads: []Thread{
			{
				Number: 1,
				Commands: []Command{
					{
						Name:       "bundle install --path vendor/bundle",
						Result:     "0",
						Output:     "Here goes long command output",
						StartTime:  mustTimeParse("2013-07-23T08:58:38Z"),
						FinishTime: mustTimeParse("2013-07-23T08:58:40Z"),
						Duration:   "00:02",
					},
					{
						Name:       "bundle exec cap -S revision=$REVISION production deploy:migrations",
						Result:     "0",
						Output:     "Here goes long command output",
						StartTime:  mustTimeParse("2013-07-23T08:58:40Z"),
						FinishTime: mustTimeParse("2013-07-23T08:59:56Z"),
						Duration:   "01:16",
					},
				},
			},
		},
		DeployInfoURL: "https://semaphoreci.com/api/v1/projects/:hash_id/servers/11/deploys/27?auth_token=:auth_token",
	}

	assert.DeepEqual(t, expected, info)
}

func TestDeploysService_Stop(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/servers/666/deploys/123/stop", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/DeploysService_Stop.json")
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

	deploy, _, err := client.Deploys.Stop("p", 666, 123)
	assert.NilError(t, err)

	expected := &Deploy{
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
	}

	assert.DeepEqual(t, expected, deploy)
}
