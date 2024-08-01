package repositories

import (
	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/djfemz/savannahTechTask/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var testGithubRepository = &models.GithubRepository{
	Name:        "test1",
	Description: "test description",
	RepoId:      2334,
	URL:         "testURL",
}

func TestSave(t *testing.T) {

	githubAuxiliaryRepository := new(mocks.GithubAuxiliaryRepository)
	githubAuxiliaryRepository.On("Save", mock.Anything).Return(testGithubRepository, nil)
	savedRepo, err := githubAuxiliaryRepository.Save(testGithubRepository)
	assert.Nil(t, err)
	assert.NotNil(t, savedRepo)
}

func TestFindByName(t *testing.T) {
	githubAuxiliaryRepository := new(mocks.GithubAuxiliaryRepository)
	githubAuxiliaryRepository.On("FindByName", mock.Anything).Return(testGithubRepository, nil)
	savedRepo, err := githubAuxiliaryRepository.FindByName("test name")
	assert.Nil(t, err)
	assert.NotNil(t, savedRepo)
}

func TestFindById(t *testing.T) {
	githubAuxiliaryRepository := new(mocks.GithubAuxiliaryRepository)
	githubAuxiliaryRepository.On("FindById", mock.Anything).Return(testGithubRepository, nil)
	savedRepo, err := githubAuxiliaryRepository.FindById(1)
	assert.Nil(t, err)
	assert.NotNil(t, savedRepo)
}
