package test

import (
	"encoding/json"
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/djfemz/savannahTechTask/api/services"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"os"
	"testing"
)

var githubMetaDataRepository = new(mocks.GithubAuxiliaryRepository)
var repoDiscoveryService = services.NewRepoDiscoveryService(services.NewGithubRepoMetadataService(githubMetaDataRepository))

func TestFetchRepoMetaData(t *testing.T) {
	var expected *dtos.GithubRepositoryResponse
	githubMetaDataRepository.On("ExistsByName", mock.Anything).Return(true)
	testGithubRepositoryMetaData := loadTestGithubRepositoryMetaData()
	githubMetaDataRepository.On("UpdateByName", mock.Anything).Return(testGithubRepositoryMetaData, nil)
	githubMetaDataRepository.On("Save", mock.Anything).Return(testGithubRepositoryMetaData, nil)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(http.MethodGet, os.Getenv("GITHUB_API_REPOSITORY_URL"), func(request *http.Request) (*http.Response, error) {
		res, err := httpmock.NewJsonResponse(http.StatusOK, testGithubRepositoryMetaData)
		return res, err
	})
	data, err := repoDiscoveryService.FetchRepoMetaData()
	expectedJsonResponse, _ := json.Marshal(testGithubRepositoryMetaData)
	err = json.Unmarshal(expectedJsonResponse, &expected)
	assert.Nil(t, err)
	assert.Equal(t, data, expected)
	assert.Nil(t, err)
	assert.NotNil(t, data)
}
