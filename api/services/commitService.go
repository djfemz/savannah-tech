package services

import (
	"github.com/djfemz/savannahTechTask/api/appErrors"
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/mappers"
	"github.com/djfemz/savannahTechTask/api/models"
	"github.com/djfemz/savannahTechTask/api/repositories"
	"log"
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

func (commitService *CommitService) GetAllCommits() (responses []*dtos.CommitResponse, err error) {
	commits, err := commitService.repository.FindAll()
	if err != nil {
		return nil, appErrors.NewCommitNotFoundError()
	}
	responses = mappers.MapToCommitResponses(commits)
	return responses, err
}

func (commitService *CommitService) GetCommitsByDateSince(since string) (response []*dtos.CommitResponse, err error) {
	sinceTime, err := time.Parse("01-02-2006", since)
	log.Println("since: ", sinceTime)
	commits, err := commitService.repository.FindAllByDateSince(&sinceTime)
	if err != nil {
		return nil, appErrors.NewCommitNotFoundError()
	}
	return mappers.MapToCommitResponses(commits), err
}

func (commitService *CommitService) GetMostRecentCommit() (*models.Commit, error) {
	return commitService.repository.FindMostRecentCommit()
}

func (commitService *CommitService) GetTopCommitAuthors(size int) ([]*dtos.AuthorResponse, error) {
	authors, err := commitService.repository.FindTopCommitAuthors(size)
	if err != nil {
		return nil, appErrors.NewAuthorNotFoundError()
	}
	authorRes := mapToAuthorResponse(authors)
	return authorRes, nil
}

func (commitService *CommitService) GetCommitsForRepo(repoName string) ([]*dtos.CommitResponse, error) {
	commits, err := commitService.repository.FindCommitsForRepoByName(repoName)
	if err != nil {
		return nil, appErrors.NewCommitNotFoundError()
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
