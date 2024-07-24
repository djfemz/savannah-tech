package test

import (
	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/djfemz/savannahTechTask/api/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

func init() {
	err := godotenv.Load("example.env")
	if err != nil {
		log.Fatal(err)
	}
}

var commitRepository = new(mocks.CommitRepository)

func TestFindLastCommitByDate(t *testing.T) {
	commitRepository.On("FindMostRecentCommit").Return(testCommits[0], nil)
	commit, err := commitRepository.FindMostRecentCommit()
	assert.Nil(t, err)
	assert.NotNil(t, commit)
}

func TestFindTopCommitAuthors(t *testing.T) {
	commitRepository.On("FindTopCommitAuthors", mock.Anything).Return(loadTestAuthors(), nil)
	authors, err := commitRepository.FindTopCommitAuthors(3)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(authors))
}

func loadTestAuthors() []*models.Author {
	return []*models.Author{
		{
			ID:       1,
			Email:    "test@author.com",
			Username: "author",
		},
		{
			ID:       2,
			Email:    "test@author1.com",
			Username: "author1",
		},
	}
}
