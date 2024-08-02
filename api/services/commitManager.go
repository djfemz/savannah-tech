package services

import (
	"github.com/djfemz/savannahTechTask/api/mappers"
	"github.com/djfemz/savannahTechTask/api/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
)

type CommitManager struct {
	*CommitService
	*RepoDiscoveryService
}

var totalCommitCount int

const (
	MAX_RECORDS_PER_PAGE = 100
)

func NewCommitManager(commitService *CommitService, repoDiscovery *RepoDiscoveryService) *CommitManager {
	return &CommitManager{commitService, repoDiscovery}
}

func (commitManager *CommitManager) FetchCommitDataFrom(since *time.Time) (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
	log.Printf("[info: start fetch commit data in %s ]", "commit manager")
	githubCommitResponses, err = commitManager.fetchAllCommits(githubCommitResponses, since)
	return githubCommitResponses, err
}

func (commitManager *CommitManager) StartJob() {
	since, _ := time.Parse("01-02-2006", os.Getenv("FETCH_DATE_SINCE"))
	log.Println(since)
	go commitManager.FetchCommitDataFrom(nil)
}

func (commitManager *CommitManager) fetchAllCommits(githubCommitResponses *[]dtos.GitHubCommitResponse, since *time.Time) (*[]dtos.GitHubCommitResponse, error) {
	counter := 1
	repoName := os.Getenv("REPO_NAME")
	repository, _ := commitManager.FindByName(repoName)

	totalCommitCount, _ = strconv.Atoi(utils.GetCommitCount())
	var responses = make([]dtos.GitHubCommitResponse, 0)
	for {
		log.Println("[INFO:]\tfetching records on page: ", counter)
		resp, err := getData(os.Getenv("GITHUB_API_COMMIT_URL"), counter, since)
		if err != nil {
			log.Println("[ERROR:]\terror fetching commit data: ", err)
		} else if resp.StatusCode == http.StatusForbidden {
			log.Printf("[info]:\tGitHub responded with a status code: %d, server sleeping off for 1 hour...", resp.StatusCode)
			time.Sleep(70 * time.Minute)
		} else if resp.StatusCode == http.StatusOK {
			log.Println("[INFO:]\tsuccess fetching...")
			githubCommitResponses, _ = extractDataInto(resp, githubCommitResponses)
			responses = append(responses, *githubCommitResponses...)
			if responses != nil && len(responses) > 0 {

				commits := mappers.MapToCommits(&responses, repository)
				if err = commitManager.repository.SaveAll(commits); err != nil {
					log.Println("[ERROR:]\terror saving data: ", err)
				}
			}
			if counter*MAX_RECORDS_PER_PAGE >= totalCommitCount {
				break
			}
			counter++
		}
	}
	log.Println("api response: ", responses)
	return &responses, nil
}
