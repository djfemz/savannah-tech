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
