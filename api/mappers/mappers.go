package mappers

import (
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/models"
)

func MapToCommits(commits *[]dtos.GitHubCommitResponse) []*models.Commit {
	var usersCommits = make([]*models.Commit, 0)
	for _, commit := range *commits {
		userCommit := models.NewCommitFromGithubCommitResponse(&commit)
		usersCommits = append(usersCommits, userCommit)
	}
	return usersCommits
}

func MapToCommitResponses(commits []*models.Commit) []*dtos.CommitResponse {
	var usersCommits = make([]*dtos.CommitResponse, 0)
	for _, commit := range commits {
		userCommit := models.NewCommitResponse(commit)
		usersCommits = append(usersCommits, userCommit)
	}
	return usersCommits
}
