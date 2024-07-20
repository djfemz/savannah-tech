package services

import (
	"github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/repositories"
)

type CommitService struct {
	repository *repositories.CommitRepository
}

func NewCommitService(repository *repositories.CommitRepository) *CommitService {
	return &CommitService{
		repository: repository,
	}
}

func (commitService CommitService) GetAllCommits() ([]*responses.CommitResponse, error) {
	return nil, nil
}

func FetchCommits() {

}
