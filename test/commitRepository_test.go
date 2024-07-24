package test

import (
	"github.com/djfemz/savannahTechTask/api/mocks"
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
	commitRepository.On("FindTopCommitAuthors", mock.Anything).Return(loadTestAuthorData(), nil)
	authors, err := commitRepository.FindTopCommitAuthors(3)
	assert.Nil(t, err)
	assert.Equal(t, len(loadTestAuthorData()), len(authors))
}
