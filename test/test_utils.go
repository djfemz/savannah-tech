package test

import (
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/utils"
	"log"
	"os"
	"time"
)

func loadTestRepositories() []*dtos.GitHubCommitResponse {
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

func GetByDate(since time.Time) (response *[]dtos.GitHubCommitResponse) {
	for _, repository := range loadTestRepositories() {
		if repository.CommitDate == since {
			*response = append(*response, *repository)
		}
	}
	return response
}
