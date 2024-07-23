package services

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/mappers"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

type CommitManager struct {
	*CommitService
}

func NewCommitManager(commitService *CommitService) *CommitManager {
	return &CommitManager{}
}

func (commitManager *CommitManager) FetchCommitDataFrom(page uint64, since time.Time) (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
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
		for counter := 1; counter < 20000; counter++ {
			go commitManager.fetch(counter)
		}
	})
	if err != nil {
		log.Println("Error creating job: ", err)
	}
	log.Println("Starting new task...")
	job.Start()
}

func (commitManager *CommitManager) fetch(counter int) {
	since, _ := time.Parse("01-02-2006", os.Getenv("FETCH_DATE_SINCE"))
	data, err := commitManager.FetchCommitDataFrom(uint64(counter), since)
	if err != nil {
		log.Println("Error fetching commits: ", err)
	}
	commits := mappers.MapToCommits(data)
	err = commitManager.repository.SaveAll(commits)
}
