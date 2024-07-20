package test

import (
	"github.com/djfemz/savannahTechTask/app/repositories"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/djfemz/savannahTechTask/app/services"
)

// TODO: replace real database instance with mock instance
func TestFetchCommitDataForRepository(t *testing.T) {
	db, _ := repositories.ConnectToDatabase()
	commitRepository := repositories.NewCommitRepository(db)
	commitService := services.NewCommitService(commitRepository)
	commits, err := commitService.GetAllCommits()
	numberOfCommitsBeforeFetch := len(commits)
	assert.Nil(t, err)
	assert.NotNil(t, commits)
	services.FetchCommits()
	commits, err = commitService.GetAllCommits()
	numberOfCommitsAfterFetch := len(commits)
	assert.GreaterOrEqual(t, numberOfCommitsAfterFetch, numberOfCommitsBeforeFetch)
}
