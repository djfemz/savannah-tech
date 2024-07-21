package test

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/mocks"
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/jarcoal/httpmock"
	"net/http"
	"os"
	"testing"
	"time"
)

var githubCommitResponse = []*dtos.GitHubCommitResponse{
	{
		Sha: "abcdefgh",
		RepoCommit: dtos.RepoCommit{
			RepoAuthor: dtos.RepoAuthor{
				Name:  "john",
				Email: "john@email.com",
				Date:  time.Now(),
			},
			Committer: dtos.Committer{
				Name:  "john",
				Email: "john@email.com",
				Date:  time.Now(),
			},
			Message: "refactored project",
		},
	},
	{
		Sha: "zyavsfe",
		RepoCommit: dtos.RepoCommit{
			RepoAuthor: dtos.RepoAuthor{
				Name:  "jane",
				Email: "jane@email.com",
				Date:  time.Now(),
			},
			Committer: dtos.Committer{
				Name:  "jane",
				Email: "jane@email.com",
				Date:  time.Now(),
			},
			Message: "initial commit",
		},
	},
}

func TestFetchCommitData(t *testing.T) {
	commitRepository := mocks.NewCommitRepository(t)
	githubAuxiliaryRepository := mocks.NewGithubAuxiliaryRepository(t)
	githubService := services.NewGithubService(services.NewCommitService(commitRepository),
		services.NewGithubRepoService(githubAuxiliaryRepository))
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), func(request *http.Request) (*http.Response, error) {
		res, err := httpmock.NewJsonResponse(http.StatusOK, githubCommitResponse)
		return res, err
	})

	githubService.FetchCommits()

}
