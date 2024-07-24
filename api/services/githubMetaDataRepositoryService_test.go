package services

import (
	"testing"

	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExistsByName(t *testing.T) {
	appRepository := mocks.NewGithubAuxiliaryRepository(t)
	appRepository.On("ExistsByName", mock.Anything).Return(true, nil)
	exists, err := appRepository.ExistsByName("test1")
	assert.True(t, exists)
	assert.Nil(t, err)
}
