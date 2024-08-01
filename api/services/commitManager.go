package services

import (
	"log"
	"os"
	"time"

	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/mappers"
)

type CommitManager struct {
	*CommitService
}

func NewCommitManager(commitService *CommitService) *CommitManager {
	return &CommitManager{commitService}
}

func (commitManager *CommitManager) FetchCommitDataFrom(since time.Time) (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
	resp, err := getData(os.Getenv("GITHUB_API_COMMIT_URL"), &since)
	if err != nil {
		return nil, err
	}
	githubCommitResponses, err = extractDataInto(resp, githubCommitResponses)
	if err != nil {
		return nil, err
	}
	return githubCommitResponses, nil
}

func (commitManager *CommitManager) StartJob() {
	go commitManager.fetch()
}

func (commitManager *CommitManager) fetch() {
	since, _ := time.Parse("01-02-2006", os.Getenv("FETCH_DATE_SINCE"))
	data, err := commitManager.FetchCommitDataFrom(since)
	if err != nil {
		log.Println("Error fetching commits: ", err)
	}
	if data != nil && len(*data) > 0 {
		commits := mappers.MapToCommits(data)
		if err = commitManager.repository.SaveAll(commits); err != nil {
			log.Println("error saving data: ", err)
		}
	}

}
