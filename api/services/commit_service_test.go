package services

import (
	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/djfemz/savannahTechTask/api/models"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testCommits = []*models.Commit{{ID: 1, RepoName: "test repo", CommitHash: "abc123", CommittedAt: time.Now()},
	{ID: 2, RepoName: "test repo 1", CommitHash: "abc1234", CommittedAt: time.Now()}}

var testAuthors = []*models.Author{{ID: 1, Name: "author 1", Email: "john@email.com"}, {ID: 1, Name: "author 1", Email: "john@email.com"}, {ID: 1, Name: "author 1", Email: "john@email.com"}}

func TestGetCommitsByDateSince(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindAllByDateSince", mock.Anything).Return(
		testCommits, nil,
	)
	commitService := NewCommitService(commitRepository)
	since := "2024-04-15T00:00:00Z"
	commits, err := commitService.GetCommitsByDateSince(since)
	assert.Nil(t, err)
	assert.NotNil(t, commits)
}

func TestGetMostRecentCommit(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindMostRecentCommit", mock.Anything).Return(
		testCommits[0], nil)
	commitService := NewCommitService(commitRepository)
	commit, err := commitService.GetMostRecentCommit()
	assert.NotNil(t, commit)
	assert.Nil(t, err)
}

func TestGetTopCommitAuthors(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindTopCommitAuthors", 3).Return(
		testAuthors, nil,
	)
	commitService := NewCommitService(commitRepository)
	authors, err := commitService.GetTopCommitAuthors(3)
	assert.NotNil(t, authors)
	assert.Nil(t, err)
}

func TestGetCommitsForRepository(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindCommitsForRepoByName", mock.Anything).Return(
		testCommits, nil,
	)
	commitService := NewCommitService(commitRepository)
	commits, err := commitService.GetCommitsForRepo("shoppersDelight")
	assert.Nil(t, err)
	assert.NotEmpty(t, commits)
}
