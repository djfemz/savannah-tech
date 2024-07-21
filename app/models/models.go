package models

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"gorm.io/gorm"
	"time"
)

type Commit struct {
	*gorm.Model
	CommitHash string `gorm:"unique"`
	Message    string
	Author     *Author `gorm:"foreignKey:CommitID"`
	Date       time.Time
	URL        string
}

type Author struct {
	*gorm.Model
	CommitID uint
	Name     string
	Email    string
	Date     time.Time
}

type AppRepository struct {
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

func NewAppRepository(response *dtos.GithubRepositoryResponse) *AppRepository {
	return &AppRepository{
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

func NewRepositoryResponse(appRepository *AppRepository) *dtos.RepositoryResponse {
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
