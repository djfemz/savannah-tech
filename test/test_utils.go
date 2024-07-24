package test

import (
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/models"
	"github.com/djfemz/savannahTechTask/api/utils"
	"log"
	"os"
	"time"
)

func loadTestGithubRepositoryMetaData() *dtos.GithubRepositoryResponse {
	return &dtos.GithubRepositoryResponse{
		ID:       23,
		NodeID:   "22345678422",
		Name:     "chromium",
		FullName: "chromium",
		Private:  false,
	}

}

func loadTestAuthorData() []*models.Author {
	return []*models.Author{
		{
			Username: "author1",
			Email:    "author1@email.com",
			Commits:  40,
		},

		{
			Username: "author2",
			Email:    "author2@email.com",
			Commits:  45,
		},
		{
			Username: "author3",
			Email:    "author3@email.com",
			Commits:  56,
		},
	}
}

func loadTestGithubCommitData() []*dtos.GitHubCommitResponse {
	var since, err = utils.GetTimeFrom(os.Getenv("ISO_TIME_FORMAT"))
	if err != nil {
		log.Println(err)
	}
	return []*dtos.GitHubCommitResponse{
		{
			Sha: "abcdefgh",
			RepoCommit: dtos.RepoCommit{
				RepoAuthor: dtos.RepoAuthor{
					Name:       "john",
					Email:      "john@email.com",
					CommitDate: *since,
				},
				Committer: dtos.Committer{
					Name:  "john",
					Email: "john@email.com",
					Date:  time.Now(),
				},
				Message: "refactored project",
			},
		},
		{
			Sha: "zyavsfe",
			RepoCommit: dtos.RepoCommit{
				RepoAuthor: dtos.RepoAuthor{
					Name:       "jane",
					Email:      "jane@email.com",
					CommitDate: time.Now(),
				},
				Committer: dtos.Committer{
					Name:  "jane",
					Email: "jane@email.com",
					Date:  time.Now(),
				},
				Message: "initial commit",
			},
		},

		{
			Sha: "zyayefe",
			RepoCommit: dtos.RepoCommit{
				RepoAuthor: dtos.RepoAuthor{
					Name:       "jane",
					Email:      "jane@email.com",
					CommitDate: time.Now(),
				},
				Committer: dtos.Committer{
					Name:  "jane",
					Email: "jane@email.com",
					Date:  time.Now(),
				},
				Message: "refactored services",
			},
		},
	}
}

func loadTestCommits() []*models.Commit {
	return []*models.Commit{
		{
			ID:      33,
			Message: "initial commit",
			Author: &models.Author{
				ID:    44,
				Email: "author@email.com",
			},
		},
		{
			ID:      35,
			Message: "refactored repo",
			Author: &models.Author{
				ID:    44,
				Email: "author1@email.com",
			},
		},
	}
}

func GetByDate(since time.Time) (response *[]dtos.GitHubCommitResponse) {
	for _, repository := range loadTestGithubCommitData() {
		if repository.CommitDate == since {
			*response = append(*response, *repository)
		}
	}
	return response
}
