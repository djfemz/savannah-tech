package services

import (
	"log"
	"os"
	"time"

	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/mappers"
	"github.com/robfig/cron/v3"
)

type CommitManager struct {
	*CommitService
}

func NewCommitManager(commitService *CommitService) *CommitManager {
	return &CommitManager{commitService}
}

func (commitManager *CommitManager) FetchPageOfCommitDataFrom(page uint64, since time.Time) (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
	resp, err := getData(os.Getenv("GITHUB_API_COMMIT_URL"), page, &since)
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
	job := cron.New()
	_, err := job.AddFunc("* * */1 * *", func() {
		go commitManager.fetch(0)
	})
	if err != nil {
		log.Println("Error creating job: ", err)
	}
	log.Println("Starting new task...")
	job.Start()
}

func (commitManager *CommitManager) fetch(page int) {
	since, _ := time.Parse("01-02-2006", os.Getenv("FETCH_DATE_SINCE"))
	data, err := commitManager.FetchPageOfCommitDataFrom(uint64(page), since)
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
