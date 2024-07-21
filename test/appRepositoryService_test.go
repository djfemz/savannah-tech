package test

import (
	"github.com/djfemz/savannahTechTask/app/mocks"
	"github.com/djfemz/savannahTechTask/app/models"
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var testGithubRepository = &models.GithubAuxiliaryRepository{
	Name:        "test1",
	Description: "test description",
	RepoId:      2334,
	URL:         "testURL",
}

func TestFetchGithubRepositoryMetaData(t *testing.T) {
	appRepository := mocks.NewGithubAuxiliaryRepository(t)
	appRepository.On("Save", mock.Anything).Return(testGithubRepository, nil)
	appRepository.On("FindByName", mock.Anything).Return(testGithubRepository, nil)
	appRepositoryService := services.NewGithubRepoService(appRepository)
	githubRepoName := "shoppersDelight"
	appRepositoryService.FetchRepositoryMetaData()
	foundRepository, err := appRepositoryService.GetRepositoryBy(githubRepoName)
	assert.Nil(t, err)
	assert.NotNil(t, foundRepository)
}
