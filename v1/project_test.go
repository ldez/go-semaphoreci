package v1

import (
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestProjectsService_Get(t *testing.T) {
	client, mux, tearDown := setupTest()
	defer tearDown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)

		content, err := ioutil.ReadFile("fixtures/ProjectsService_Get.json")
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

	projects, _, err := client.Projects.Get()
	assert.NilError(t, err)

	expected := []Project{
		{
			ID:        61,
			HashID:    "3f1004b8343faabda63d441734526c854380ab89",
			Name:      "testapp-sphinx",
			Owner:     "renderedtext",
			HTMLURL:   "https://semaphoreci.com/darkofabijan/testapp-sphinx",
			CreatedAt: mustTimeParse("2012-09-04T11:53:22Z"),
			UpdatedAt: mustTimeParse("2012-09-04T12:01:17Z"),
			Branches: []ProjectBranch{
				{
					Name:         "master",
					URL:          "https://semaphoreci.com/projects/61/branches/85",
					StatusURL:    "https://semaphoreci.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/status?auth_token=:auth_token",
					HistoryURL:   "https://semaphoreci.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85?auth_token=:auth_token",
					ProjectName:  "testapp-sphinx",
					BuildURL:     "https://semaphoreci.com/projects/61/branches/85/builds/1",
					BuildInfoURL: "https://semaphoreci.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=:auth_token",
					BuildNumber:  1,
					Result:       "passed",
					StartedAt:    mustTimeParse("2012-09-04T11:55:07Z"),
					FinishedAt:   mustTimeParse("2012-09-04T12:01:16Z"),
				},
			},
			Servers: []ServerStatus{
				{
					ProjectName:  "heroku-deploy-test",
					ServerName:   "server-heroku-master-automatic-1",
					Number:       1,
					Result:       "passed",
					CreatedAt:    mustTimeParse("2013-07-19T14:57:18+02:00"),
					UpdatedAt:    mustTimeParse("2013-07-19T14:58:49+02:00"),
					StartedAt:    mustTimeParse("2013-07-19T14:57:24+02:00"),
					FinishedAt:   mustTimeParse("2013-07-19T14:58:49+02:00"),
					HTMLURL:      "https://semaphoreci.com/projects/2420/servers/80/deploys/1",
					DeployURL:    "https://semaphoreci.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/servers/80/deploys/1?auth_token=:auth_token",
					DeployLogURL: "https://semaphoreci.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/servers/80/deploys/1/log?auth_token=:auth_token",
					BuildURL:     "https://semaphoreci.com/api/v1/projects/dccc4ab03f2b47167cac51553d58e468c6750c2c/58394/builds/5?auth_token=:auth_token",
					BuildHTMLURL: "https://semaphoreci.com/projects/2420/branches/58394/builds/5",
					Commit: &Commit{
						ID:          "43ddb7516ecc743f0563abd7418f0bd3617348c4",
						URL:         "https://github.com/rastasheep/heroku-deploy-test/commit/43ddb7516ecc743f0563abd7418f0bd3617348c4",
						AuthorName:  "Aleksandar Diklic",
						AuthorEmail: "rastasheep3@gmail.com",
						Message:     "One more time",
						Timestamp:   mustTimeParse("2013-07-19T14:56:25+02:00"),
					},
				},
			},
		},
		{
			ID:        63,
			HashID:    "649e584dc507ca4b73e1374d3125ef0b567a949c",
			Name:      "testapp-mongodb",
			Owner:     "renderedtext",
			HTMLURL:   "https://semaphoreci.com/darkofabijan/testapp-mongodb",
			CreatedAt: mustTimeParse("2012-09-14T10:53:38Z"),
			UpdatedAt: mustTimeParse("2012-09-14T11:16:51Z"),
			Branches: []ProjectBranch{
				{
					Name:         "mongoid3",
					URL:          "https://semaphoreci.com/projects/63/branches/89",
					StatusURL:    "https://semaphoreci.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/status?auth_token=:auth_token",
					HistoryURL:   "https://semaphoreci.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85?auth_token=:auth_token",
					ProjectName:  "testapp-mongodb",
					BuildURL:     "https://semaphoreci.com/projects/63/branches/89/builds/3",
					BuildInfoURL: "https://semaphoreci.com/api/v1/projects/3f1004b8343faabda63d441734526c854380ab89/85/builds/1?auth_token=:auth_token",
					BuildNumber:  3,
					Result:       "passed",
					StartedAt:    mustTimeParse("2012-09-14T11:11:39Z"),
					FinishedAt:   mustTimeParse("2012-09-14T11:16:51Z"),
				},
			},
			Servers: []ServerStatus{},
		},
	}
	assert.DeepEqual(t, expected, projects)
}
