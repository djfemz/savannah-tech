package test

import (
	"github.com/djfemz/savannahTechTask/app/repositories"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"

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

func TestGetCommitsByDateSince(t *testing.T) {
	db, err := repositories.ConnectToDatabase()
	assert.Nil(t, err)
	assert.NotNil(t, db)
	commitRepository := repositories.NewCommitRepository(db)
	commitService := services.NewCommitService(commitRepository)
	since, err := time.Parse(time.RFC3339, "2024-04-15T00:00:00Z")
	log.Println("since: ", since)
	commits, err := commitService.GetCommitsByDateSince(since)
	assert.Nil(t, err)
	log.Println("commits: ", commits)
	assert.NotNil(t, commits)
}
