package services

import (
	"github.com/djfemz/savannahTechTask/api/utils"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/djfemz/savannahTechTask/api/mocks"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchCommitData(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	githubMetaDataRepo := new(mocks.GithubAuxiliaryRepository)
	commitRepository.On("SaveAll", mock.Anything).Return(nil)
	commitRepository.On("FindMostRecentCommit").Return(utils.LoadTestCommits()[0], nil)
	commitRepository.On("CountCommits").Return(int64(3), nil)
	githubMetaDataRepo.On("FindByName", mock.Anything).Return(utils.GetRepoMetaData(), nil)
	commitMonitorService := NewCommitMonitorService(NewCommitManager(NewCommitService(commitRepository),
		NewRepoDiscoveryService(NewGithubRepoMetadataService(githubMetaDataRepo))))
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), func(request *http.Request) (*http.Response, error) {
		res, err := httpmock.NewJsonResponse(http.StatusOK, utils.LoadTestGithubCommitData())
		return res, err
	})
	data, err := commitMonitorService.FetchCommitData()
	time.Sleep(5 * time.Second)
	assert.Nil(t, err)
	assert.NotEmpty(t, data)
}
