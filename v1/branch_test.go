package v1

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestBranchesService_GetByProject(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/branches", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/BranchesService_GetByProject.json")
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

	branches, _, err := client.Branch.GetByProject("p")
	assert.NilError(t, err)

	expected := []Branch{
		{
			ID:        1324,
			Name:      "new-build-page",
			BranchURL: "https://semaphoreci.com/api/v1/projects/:hash_id/1324/status?auth_token=:auth_token",
		},
		{
			ID:        1120,
			Name:      "development",
			BranchURL: "https://semaphoreci.com/api/v1/projects/:hash_id/1120/status?auth_token=:auth_token",
		},
		{
			ID:        987,
			Name:      "branches_api",
			BranchURL: "https://semaphoreci.com/api/v1/projects/:hash_id/987/status?auth_token=:auth_token",
		},
	}

	assert.DeepEqual(t, expected, branches)
}

func TestBranchesService_GetStatus(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666/status", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/BranchesService_GetStatus.json")
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

	statuses, err := client.Branch.GetStatus("p", 666)
	assert.NilError(t, err)

	expected := &BranchStatus{
		BranchName:       "gem_updates",
		BranchURL:        "https://semaphoreci.com/projects/44/branches/50",
		BranchStatusURL:  "https://semaphoreci.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/status?auth_token=:auth_token",
		BranchHistoryURL: "https://semaphoreci.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89?auth_token=:auth_token",
		ProjectName:      "base-app",
		BuildURL:         "https://semaphoreci.com/projects/44/branches/50/builds/15",
		BuildInfoURL:     "https://semaphoreci.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/builds/1?auth_token=:auth_token",
		BuildNumber:      15,
		Result:           "passed",
		StartedAt:        mustTimeParse("2012-07-09T15:23:53Z"),
		FinishedAt:       mustTimeParse("2012-07-09T15:30:16Z"),
		Commit: Commit{
			ID:          "dc395381e650f3bac18457909880829fc20e34ba",
			URL:         "https://github.com/renderedtext/base-app/commit/dc395381e650f3bac18457909880829fc20e34ba",
			AuthorName:  "Vladimir Saric",
			AuthorEmail: "vladimir@renderedtext.com",
			Message:     "Update 'shoulda' gem.",
			Timestamp:   mustTimeParse("2012-07-04T18:14:08Z"),
		},
	}
	assert.DeepEqual(t, expected, statuses)
}

func TestBranchesService_GetHistory(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects/p/666", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/BranchesService_GetHistory.json")
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

	history, _, err := client.Branch.GetHistory("p", 666, &BranchHistoryOptions{Page: 1})
	assert.NilError(t, err)

	expected := &BranchHistory{
		BranchName:       "gem_updates",
		BranchURL:        "https://semaphoreci.com/projects/44/branches/50",
		BranchStatusURL:  "https://semaphoreci.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89/status?auth_token=:auth_token",
		BranchHistoryURL: "https://semaphoreci.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/89?auth_token=:auth_token",
		ProjectName:      "base-app",
		Builds: []Build{
			{
				BuildURL:     "https://semaphoreci.com/projects/57/branches/80/builds/46",
				BuildInfoURL: "https://semaphoreci.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/80/builds/1?auth_token=:auth_token",
				BuildNumber:  46,
				Result:       "failed",
				StartedAt:    mustTimeParse("2012-10-02T15:01:41Z"),
				FinishedAt:   mustTimeParse("2012-10-02T15:03:53Z"),
				Commit: Commit{
					ID:          "a31d32d5de89613369f934eb7d30fbeb08883334",
					URL:         "https://github.com/renderedtext/base-app/commit/a31d32d5de89613369f934eb7d30fbeb08883334",
					AuthorName:  "Vladimir Saric",
					AuthorEmail: "vladimir@renderedtext.com",
					Message:     "Update 'shoulda' gem.",
					Timestamp:   mustTimeParse("2012-10-02T07:00:14Z"),
				},
			},
			{
				BuildURL:     "https://semaphoreci.com/projects/57/branches/80/builds/45",
				BuildInfoURL: "https://semaphoreci.com/api/v1/projects/649e584dc507ca4b73e1374d3125ef0b567a949c/80/builds/1?auth_token=:auth_token",
				BuildNumber:  45,
				Result:       "passed",
				StartedAt:    mustTimeParse("2012-10-02T14:47:06Z"),
				FinishedAt:   mustTimeParse("2012-10-02T14:51:35Z"),
				Commit: Commit{
					ID:          "73fce130ad23f265add5d55ee1da1c23b38f85a4",
					URL:         "https://github.com/renderedtext/base-app/commit/73fce130ad23f265add5d55ee1da1c23b38f85a4",
					AuthorName:  "Marko Anastasov",
					AuthorEmail: "marko@renderedtext.com",
					Message:     "Update 'factory_girl_rails' gem and use new short FactoryGirl syntax.",
					Timestamp:   mustTimeParse("2012-10-02T07:00:14Z"),
				},
			},
		},
	}
	assert.DeepEqual(t, expected, history)
}
