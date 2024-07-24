package repositories

import (
	"github.com/djfemz/savannahTechTask/api/mocks"
	"github.com/djfemz/savannahTechTask/api/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var commitRepository = new(mocks.CommitRepository)

func TestFindLastCommitByDate(t *testing.T) {
	commitRepository.On("FindMostRecentCommit").Return(utils.LoadTestCommits()[0], nil)
	commit, err := commitRepository.FindMostRecentCommit()
	assert.Nil(t, err)
	assert.NotNil(t, commit)
}

func TestFindTopCommitAuthors(t *testing.T) {
	commitRepository.On("FindTopCommitAuthors", mock.Anything).Return(utils.LoadTestAuthorData(), nil)
	authors, err := commitRepository.FindTopCommitAuthors(3)
	assert.Nil(t, err)
	assert.Equal(t, len(utils.LoadTestAuthorData()), len(authors))
}
