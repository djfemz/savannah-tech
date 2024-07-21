package test

import (
	"github.com/djfemz/savannahTechTask/app/repositories"
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchGithubRepositoryMetaData(t *testing.T) {
	db, _ := repositories.ConnectToDatabase()
	appRepository := repositories.NewGithubRepoRepository(db)
	appRepositoryService := services.NewGithubRepoService(appRepository)
	githubRepoName := "shoppersDelight"
	appRepositoryService.FetchRepositoryMetaData()
	foundRepository, err := appRepositoryService.GetRepository(githubRepoName)
	assert.Nil(t, err)
	assert.NotNil(t, foundRepository)
}
