package test

import (
	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

// TODO: complete test
func TestFindTopCommitAuthors(t *testing.T) {
	commitRepository.On("FindMostRecentCommit").Return(testCommits[0], nil)
}
