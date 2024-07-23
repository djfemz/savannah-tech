package test

import (
	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/djfemz/savannahTechTask/api/models"
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

func TestExistsByName(t *testing.T) {
	appRepository := mocks.NewGithubAuxiliaryRepository(t)
	appRepository.On("ExistsByName", mock.Anything).Return(true, nil)
	exists, err := appRepository.ExistsByName("test1")
	assert.True(t, exists)
	assert.Nil(t, err)
}
