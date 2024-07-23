package test

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"time"
)

var githubCommitResponse = []*dtos.GitHubCommitResponse{
	{
		Sha: "abcdefgh",
		RepoCommit: dtos.RepoCommit{
			RepoAuthor: dtos.RepoAuthor{
				Name:  "john",
				Email: "john@email.com",
				Date:  time.Now(),
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
				Name:  "jane",
				Email: "jane@email.com",
				Date:  time.Now(),
			},
			Committer: dtos.Committer{
				Name:  "jane",
				Email: "jane@email.com",
				Date:  time.Now(),
			},
			Message: "initial commit",
		},
	},
}
