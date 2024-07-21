package services

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/mappers"
	"github.com/djfemz/savannahTechTask/app/models"
	"github.com/djfemz/savannahTechTask/app/repositories"
	"time"
)

type CommitService struct {
	repository repositories.CommitRepository
}

func NewCommitService(repository repositories.CommitRepository) *CommitService {
	return &CommitService{
		repository: repository,
	}
}

// TODO: Add Caching Strategies
func (commitService *CommitService) GetAllCommits() (responses []*dtos.CommitResponse, err error) {
	commits, err := commitService.repository.FindAll()
	if err != nil {
		return nil, err
	}
	responses = mappers.MapToCommitResponses(commits)
	return responses, err
}

func (commitService *CommitService) GetCommitsByDateSince(since time.Time) (response []*dtos.CommitResponse, err error) {
	commits, err := commitService.repository.FindAllByDateSince(&since)
	if err != nil {
		return nil, err
	}
	return mappers.MapToCommitResponses(commits), err
}

func (commitService *CommitService) GetMostRecentCommit() (*models.Commit, error) {
	return commitService.repository.FindMostRecentCommit()
}

func (commitService *CommitService) GetTopCommitAuthors(size int) ([]*dtos.AuthorResponse, error) {
	authors, err := commitService.repository.FindTopCommitAuthors(size)
	if err != nil {
		return nil, err
	}
	authorRes := mapToAuthorResponse(authors)
	return authorRes, nil
}

func (commitService *CommitService) GetCommitsForRepo(repoName string) ([]*dtos.CommitResponse, error) {
	commits, err := commitService.repository.FindCommitsForRepoByName(repoName)
	if err != nil {
		return nil, err
	}
	commitRes := mappers.MapToCommitResponses(commits)
	return commitRes, nil
}

func mapToAuthorResponse(authors []*models.Author) []*dtos.AuthorResponse {
	authorRes := make([]*dtos.AuthorResponse, 0)
	for _, author := range authors {
		authorRes = append(authorRes, models.NewAuthorResponse(author))
	}
	return authorRes

}
