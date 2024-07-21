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
	appRepository := new(mocks.GithubAuxiliaryRepository)
	commitRepository := new(mocks.CommitRepository)
	appRepository.On("Save", mock.Anything).Return(testGithubRepository, nil)
	appRepository.On("FindByName", mock.Anything).Return(testGithubRepository, nil)
	appRepository.On("ExistsByName", mock.Anything).Return(true, nil)
	appRepository.On("UpdateByName", mock.Anything, mock.Anything).Return(testGithubRepository, nil)
	appRepositoryService := services.NewGithubService(services.NewCommitService(commitRepository), services.NewGithubRepoService(appRepository))
	githubRepoName := "shoppersDelight"
	appRepositoryService.FetchRepositoryMetaData()
	foundRepository, err := appRepositoryService.GetRepositoryBy(githubRepoName)
	assert.Nil(t, err)
	assert.NotNil(t, foundRepository)
}

func TestExistsByName(t *testing.T) {
	appRepository := mocks.NewGithubAuxiliaryRepository(t)
	appRepository.On("ExistsByName", mock.Anything).Return(true, nil)
	exists, err := appRepository.ExistsByName("test1")
	assert.True(t, exists)
	assert.Nil(t, err)
}
