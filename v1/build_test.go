package v1

import (
	"net/http"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestBuildsService_GetInformation(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666/builds/123", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/BuildsService_GetInformation.json")
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

	info, err := client.Builds.GetInformation("p", 666, 123)
	assert.NilError(t, err)

	expected := &BuildInformation{
		Commits: []Commit{
			{
				ID:          "ce0d543b875884f09cf1e287fb303fb91a9e28a0",
				URL:         "https://github.com/renderedtext/base-app/commit/ce0d543b875884f09cf1e287fb303fb91a9e28a0",
				AuthorName:  "Marko Anastasov",
				AuthorEmail: "marko@renderedtext.com",
				Message:     "Upgrade shoulda 1.1.6 -> 1.2.1",
				Timestamp:   mustTimeParse("2014-05-16T15:38:49+02:00"),
			},
		},
		ProjectName: "base-app",
		BranchName:  "master",
		Number:      44,
		Result:      "passed",
		CreatedAt:   mustTimeParse("2014-06-19T09:37:18+02:00"),
		UpdatedAt:   mustTimeParse("2014-06-19T09:39:42+02:00"),
		StartedAt:   mustTimeParse("2014-06-19T09:37:26+02:00"),
		FinishedAt:  mustTimeParse("2014-06-19T09:39:42+02:00"),
		HTMLURL:     "https://semaphoreci.com/renderedtext/base-app/branches/master/builds/44",
		BuildLogURL: "https://semaphoreci.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/44/log?auth_token=:auth_token",
	}

	assert.DeepEqual(t, expected, info)
}

func TestBuildsService_GetLog(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666/builds/123/log", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := os.ReadFile("fixtures/BuildsService_GetLog.json")
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

	buildLog, err := client.Builds.GetLog("p", 666, 123)
	assert.NilError(t, err)

	expected := &BuildLog{
		Threads: []Thread{
			{
				Number: 1,
				Commands: []Command{
					{
						Name:       "bundle install --deployment --path vendor/bundle",
						Result:     "0",
						Output:     "Fetching source index from http://rubygems.org/ Installing rake (10.0.3) Installing i18n (0.6.1) Installing multi_json (1.6.1) Installing activesupport (3.2.11) Installing builder (3.0.4) Installing activemodel (3.2.11) Installing erubis (2.7.0) Installing journey (1.0.4) Installing rack (1.4.5) Installing rack-cache (1.2) Installing rack-test (0.6.2) Installing hike (1.2.1) Installing tilt (1.3.3) Installing sprockets (2.2.2) Installing actionpack (3.2.11) Installing mime-types (1.21) Installing polyglot (0.3.3) Installing treetop (1.4.12) Installing mail (2.4.4) Installing actionmailer (3.2.11) Installing arel (3.0.2) Installing tzinfo (0.3.35) Installing activerecord (3.2.11) Installing activeresource (3.2.11) Installing addressable (2.3.3) Installing kaminari (0.14.1) Installing polyamorous (0.5.0) Installing meta_search (1.1.3) Using bundler (1.3.0) Installing rack-ssl (1.3.2) Installing json (1.7.7)",
						StartTime:  mustTimeParse("2013-03-12T09:24:30Z"),
						FinishTime: mustTimeParse("2013-03-12T09:24:37Z"),
						Duration:   "00:45",
					},
					{
						Name:       "gem --version",
						Result:     "0",
						Output:     "1.8.23",
						StartTime:  mustTimeParse("2013-03-12T09:25:30Z"),
						FinishTime: mustTimeParse("2013-03-12T09:25:37Z"),
						Duration:   "00:00",
					},
				},
			},
			{
				Number: 2,
				Commands: []Command{
					{
						Name:       "bundle exec rake db:test:prepare",
						Result:     "0",
						Output:     "",
						StartTime:  mustTimeParse("2013-03-12T09:24:37Z"),
						FinishTime: mustTimeParse("2013-03-12T09:24:44Z"),
						Duration:   "00:07",
					},
				},
			},
		},
		BuildInfoURL: "https://semaphoreci.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=:auth_token",
	}

	assert.DeepEqual(t, expected, buildLog)
}

func TestBuildsService_RebuildLastRevision(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666/build", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)

		content, err := os.ReadFile("fixtures/BuildsService_RebuildLastRevision.json")
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

	info, err := client.Builds.RebuildLastRevision("p", 666)
	assert.NilError(t, err)

	expected := &BuildInformation{
		Commits: []Commit{
			{
				ID:          "ee89ebaaaeasdaasdasdqwewlweqlqwleqlwe",
				URL:         "https://github.com/renderedtext/semaphore/commit/dasadsdasadsdasadsdsaasdasdasd",
				AuthorName:  "Rastko Jokic",
				AuthorEmail: "rastko@renderedtext.com",
				Message:     "Add cucumber_in_groups",
				Timestamp:   mustTimeParse("2013-12-25T11:57:44+01:00"),
			},
			{
				ID:          "2a6e8df4llllll11427e1asdasl97506ffac15",
				URL:         "https://github.com/renderedtext/semaphore/commit/2a6e8dasddasdsasadaf69753d5d06ffac15",
				AuthorName:  "Marko Anastasov",
				AuthorEmail: "marko@renderedtext.com",
				Message:     "Merge pull request #410 from renderedtext/rj/cucumber-groups\n\nAdd cucumber_in_groups",
				Timestamp:   mustTimeParse("2013-12-25T12:31:07+01:00"),
			},
		},
		ProjectName:  "semaphore",
		BranchName:   "development",
		Number:       1,
		Result:       "",
		CreatedAt:    mustTimeParse("2013-12-25T14:56:43+01:00"),
		UpdatedAt:    mustTimeParse("2013-12-25T14:56:44+01:00"),
		StartedAt:    nil,
		FinishedAt:   nil,
		HTMLURL:      "https://semaphoreci.com/darkofabijan/semaphore/branches/development/builds/1",
		BuildLogURL:  "https://semaphoreci.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1/log?auth_token=:auth_token",
		BuildInfoURL: "https://semaphoreci.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/1?auth_token=:auth_token",
	}

	assert.DeepEqual(t, expected, info)
}

func TestBuildsService_Launch(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666/build", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)

		content, err := os.ReadFile("fixtures/BuildsService_Launch.json")
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

	info, err := client.Builds.Launch("p", 666, "c")
	assert.NilError(t, err)

	expected := &BuildInformation{
		Commits: []Commit{
			{
				ID:          "9d4a35a78942d52ddc88d6e75dbde44d1ba4fa50",
				URL:         "https://github.com/renderedtext/base-app/commit/9d4a35a78942d52ddc88d6e75dbde44d1ba4fa50",
				AuthorName:  "Marko Anastasov",
				AuthorEmail: "marko@renderedtext.com",
				Message:     "Update cucumber-rails",
				Timestamp:   mustTimeParse("2014-01-31T13:02:27+01:00"),
			},
		},
		ProjectName:  "base-app",
		BranchName:   "master",
		Number:       45,
		Result:       "",
		CreatedAt:    mustTimeParse("2014-08-18T17:04:52+02:00"),
		UpdatedAt:    mustTimeParse("2014-08-18T17:04:53+02:00"),
		StartedAt:    nil,
		FinishedAt:   nil,
		HTMLURL:      "https://semaphoreci.com/renderedtext/base-app/branches/master/builds/45",
		BuildLogURL:  "https://semaphoreci.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/45/log?auth_token=:auth_token",
		BuildInfoURL: "https://semaphoreci.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/45?auth_token=:auth_token",
	}

	assert.DeepEqual(t, expected, info)
}

func TestBuildsService_Stop(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666/builds/123/stop", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)

		content, err := os.ReadFile("fixtures/BuildsService_Stop.json")
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

	info, err := client.Builds.Stop("p", 666, 123)
	assert.NilError(t, err)

	expected := &BuildInformation{
		Commits: []Commit{
			{
				ID:          "222f0123418545f21234184a4725fe16asfa125123",
				URL:         "https://github.com/renderedtext/semaphore/commit/222f0123418545f21234184a4725fe16asfa125123",
				AuthorName:  "Marko Anastasov",
				AuthorEmail: "marko@renderedtext.com",
				Message:     "Merge branch 'staging'",
				Timestamp:   mustTimeParse("2013-07-22T17:52:27+02:00"),
			},
		},
		ProjectName:  "base-app",
		BranchName:   "master",
		Number:       27,
		Result:       "stopped",
		CreatedAt:    mustTimeParse("2014-08-18T17:04:52+02:00"),
		UpdatedAt:    mustTimeParse("2014-08-18T17:04:53+02:00"),
		StartedAt:    mustTimeParse("2014-08-18T17:04:59+02:00"),
		FinishedAt:   nil,
		HTMLURL:      "https://semaphoreci.com/projects/1/branches/29803/builds/27",
		BuildLogURL:  "https://semaphoreci.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/27/log?auth_token=:auth_token",
		BuildInfoURL: "https://semaphoreci.com/api/v1/projects/73c4b979-0a40-49db-b10e-571d41e10d9a/133529/builds/27?auth_token=:auth_token",
	}

	assert.DeepEqual(t, expected, info)
}

func TestBuildsService_DeployFromBuild(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666/builds/123/deploy/789", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodPost)

		content, err := os.ReadFile("fixtures/BuildsService_DeployFromBuild.json")
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

	serverStatus, err := client.Builds.DeployFromBuild("p", 666, 123, 789)
	assert.NilError(t, err)

	expected := &ServerStatus{
		ProjectName:  "semaphore",
		ServerName:   "production",
		Number:       27,
		CreatedAt:    mustTimeParse("2013-07-23T10:57:42+02:00"),
		UpdatedAt:    mustTimeParse("2013-07-23T11:00:41+02:00"),
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

	assert.DeepEqual(t, expected, serverStatus)
}
