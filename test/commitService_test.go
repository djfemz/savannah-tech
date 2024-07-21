package test

import (
	"github.com/djfemz/savannahTechTask/app/mocks"
	"github.com/djfemz/savannahTechTask/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/djfemz/savannahTechTask/app/services"
)

var testCommits = []*models.Commit{{ID: 1, RepoName: "test repo", CommitHash: "abc123"},
	{ID: 2, RepoName: "test repo 1", CommitHash: "abc1234"}}

var testAuthors = []*models.Author{{ID: 1, Name: "author 1", Email: "john@email.com"}, {ID: 1, Name: "author 1", Email: "john@email.com"}, {ID: 1, Name: "author 1", Email: "john@email.com"}}

func TestFetchCommitDataForRepository(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)

	commitRepository.On("FindAll", mock.Anything).Return(
		testCommits,
		nil,
	)
	commitRepository.On("SaveAll", testCommits).Return(nil)
	commitService := services.NewCommitService(commitRepository)
	commits, err := commitService.GetAllCommits()
	numberOfCommitsBeforeFetch := len(commits)
	assert.Nil(t, err)
	assert.NotNil(t, commits)
	//services.FetchCommits()
	commits, err = commitService.GetAllCommits()
	numberOfCommitsAfterFetch := len(commits)
	assert.GreaterOrEqual(t, numberOfCommitsAfterFetch, numberOfCommitsBeforeFetch)
}

func TestGetCommitsByDateSince(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindAllByDateSince", mock.Anything).Return(
		testCommits, nil,
	)
	commitService := services.NewCommitService(commitRepository)
	since, err := time.Parse(time.RFC3339, "2024-04-15T00:00:00Z")
	commits, err := commitService.GetCommitsByDateSince(since)
	assert.Nil(t, err)
	assert.NotNil(t, commits)
}

func TestGetMostRecentCommit(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindMostRecentCommit", mock.Anything).Return(
		testAuthors, nil)
	commitService := services.NewCommitService(commitRepository)
	commit, err := commitService.GetMostRecentCommit()
	assert.NotNil(t, commit)
	assert.Nil(t, err)
}

func TestGetTopCommitAuthors(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindTopCommitAuthors", 3).Return(
		testAuthors, nil,
	)
	commitService := services.NewCommitService(commitRepository)
	authors, err := commitService.GetTopCommitAuthors(3)
	assert.NotNil(t, authors)
	assert.Nil(t, err)
}

func TestGetCommitsForRepository(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindCommitsForRepoByName", mock.Anything).Return(
		testCommits, nil,
	)
	commitService := services.NewCommitService(commitRepository)
	commits, err := commitService.GetCommitsForRepo("shoppersDelight")
	assert.Nil(t, err)
	assert.NotEmpty(t, commits)
}
