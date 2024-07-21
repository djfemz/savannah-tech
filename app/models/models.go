package models

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"gorm.io/gorm"
	"os"
	"time"
)

type Commit struct {
	ID uint `gorm:"primaryKey"`
	*gorm.Model
	RepoName   string
	CommitHash string `gorm:"unique"`
	Message    string
	Author     *Author
	Date       time.Time
	URL        string
}

type Author struct {
	ID uint
	*gorm.Model
	Username string
	Name     string
	Email    string
	CommitID uint
	Date     time.Time
	Commits  uint
}

type GithubAuxiliaryRepository struct {
	ID             uint
	Name           string
	Description    string
	RepoId         uint `gorm:"unique"`
	URL            string
	Language       string
	ForkCount      int
	StarCount      int
	OpenIssueCount int
	WatcherCount   int
	DateCreated    time.Time
	UpdatedDate    time.Time
}

func NewCommitFromGithubCommitResponse(response *dtos.GitHubCommitResponse) *Commit {
	return &Commit{
		RepoName:   os.Getenv("REPO_NAME"),
		CommitHash: response.Sha,
		Message:    response.RepoCommit.Message,
		URL:        response.RepoCommit.URL,
		Date:       response.RepoCommit.Committer.Date,
		Author: &Author{
			Name:     response.RepoCommit.RepoAuthor.Name,
			Email:    response.RepoCommit.RepoAuthor.Email,
			Username: response.Login,
		},
	}
}

func NewCommitResponse(commit *Commit) (commitResponse *dtos.CommitResponse) {
	commitResponse = &dtos.CommitResponse{
		ID:      commit.ID,
		Date:    commit.Date,
		Message: commit.Message,
		URL:     commit.URL,
	}
	if commit.Author != nil {
		commitResponse.AuthorEmail = commit.Author.Email
		commitResponse.Author = commit.Author.Name
	}
	return commitResponse
}

func NewGithubAuxiliaryRepository(response *dtos.GithubRepositoryResponse) *GithubAuxiliaryRepository {
	return &GithubAuxiliaryRepository{
		Name:           response.Name,
		Description:    response.Description,
		URL:            response.URL,
		Language:       response.Language,
		ForkCount:      response.ForksCount,
		StarCount:      response.StargazersCount,
		OpenIssueCount: response.OpenIssuesCount,
		WatcherCount:   response.WatchersCount,
		DateCreated:    response.CreatedAt,
		UpdatedDate:    response.UpdatedAt,
	}
}

func NewRepositoryResponse(appRepository *GithubAuxiliaryRepository) *dtos.RepositoryResponse {
	return &dtos.RepositoryResponse{
		Name:           appRepository.Name,
		RepoId:         appRepository.RepoId,
		URL:            appRepository.URL,
		Description:    appRepository.Description,
		Language:       appRepository.Language,
		ForkCount:      appRepository.ForkCount,
		StarCount:      appRepository.StarCount,
		OpenIssueCount: appRepository.OpenIssueCount,
		WatcherCount:   appRepository.WatcherCount,
		DateCreated:    appRepository.DateCreated,
		UpdatedDate:    appRepository.UpdatedDate,
	}

}

func NewAuthorResponse(author *Author) *dtos.AuthorResponse {
	return &dtos.AuthorResponse{
		Username: author.Username,
		Email:    author.Email,
		Commits:  author.Commits,
	}
}
