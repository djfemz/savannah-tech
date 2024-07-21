package test

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/mocks"
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
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
	commitRepository := new(mocks.CommitRepository)
	githubAuxiliaryRepository := new(mocks.GithubAuxiliaryRepository)
	githubService := services.NewGithubService(services.NewCommitService(commitRepository),
		services.NewGithubRepoService(githubAuxiliaryRepository))
	commitRepository.On("FindMostRecentCommit").Return(testCommits[0], nil)
	commitRepository.On("SaveAll", mock.Anything).Return(nil)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), func(request *http.Request) (*http.Response, error) {
		res, err := httpmock.NewJsonResponse(http.StatusOK, githubCommitResponse)
		log.Println("res--->", res)
		return res, err
	})

	commits, err := githubService.FetchCommits()
	assert.NotEmpty(t, commits)
	assert.Nil(t, err)

}
