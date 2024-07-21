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
	Author     *Author `gorm:"foreignKey:CommitID"`
	Date       time.Time
	URL        string
}

type Author struct {
	ID uint
	*gorm.Model
	CommitID uint
	Name     string
	Email    string
	Date     time.Time
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
			Name:  response.RepoCommit.Author.Name,
			Email: response.RepoCommit.Author.Email,
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

func NewAppRepository(response *dtos.GithubRepositoryResponse) *GithubAuxiliaryRepository {
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
