package services

import (
	"fmt"
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
var doneChannel *chan bool
var repoName string

const (
	MAX_RECORDS_PER_PAGE        = 100
	MINIMUM_ALLOWED_PAGE_NUMBER = 1
	COMMITS_ENDPOINT            = "/commits"
)

func init() {

}

func NewCommitManager(commitService *CommitService, repoDiscovery *RepoDiscoveryService, logger *log.Logger) *CommitManager {
	repoName = os.Getenv("REPO_NAME")
	return &CommitManager{commitService, repoDiscovery}
}

func (commitManager *CommitManager) FetchCommitDataFrom(since *time.Time) (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
	log.Printf("[info: start pullCommitDataFromGithub commit data in %s ]", "commit manager")
	githubCommitResponses, err = commitManager.fetchAllCommits(githubCommitResponses, since)
	return githubCommitResponses, err
}

func (commitManager *CommitManager) StartJob(ch *chan bool) {
	doneChannel = ch
	go commitManager.FetchCommitDataFrom(nil)
}

func (commitManager *CommitManager) fetchAllCommits(githubCommitResponses *[]dtos.GitHubCommitResponse, since *time.Time) (*[]dtos.GitHubCommitResponse, error) {
	counter, _ := commitManager.CommitService.repository.CountCommits()
	counter = counter / int64(MAX_RECORDS_PER_PAGE)

	repository, _ := commitManager.GithubAuxiliaryRepository.FindByName(repoName)

	totalCommitCount, _ = strconv.Atoi(utils.GetCommitCount())
	url := fmt.Sprintf("%s%s", os.Getenv("GITHUB_API_BASE_URL"), COMMITS_ENDPOINT)
	for {
		log.Println("[INFO:]\tfetching records on page: ", counter)
		resp, err := getData(url, counter, since)
		if err != nil {
			log.Println("[ERROR:]\terror fetching commit data: ", err)
		} else if resp.StatusCode == http.StatusForbidden && resp.Header.Get("x-ratelimit-remaining") != "0" {

			sleep(resp)
		} else if resp.StatusCode == http.StatusOK {
			log.Println("[INFO:]\tsuccess fetching page: ", counter)
			githubCommitResponses, _ = extractDataInto(resp, githubCommitResponses)
			if githubCommitResponses != nil && len(*githubCommitResponses) > 0 {
				commits := mappers.MapToCommits(githubCommitResponses, repository)
				if err = commitManager.repository.SaveAll(commits); err != nil {
					log.Println("[ERROR:]\terror saving data: ", err)
				}
			}
			isDoneFetchingCommitData := (counter * int64(MAX_RECORDS_PER_PAGE)) >= int64(totalCommitCount)
			if isDoneFetchingCommitData {
				break
			}
			counter++
		}
	}
	if doneChannel == nil {
		return githubCommitResponses, nil
	} else {
		ch := make(chan bool)
		doneChannel = &ch
		*doneChannel <- true
		close(*doneChannel)
	}
	return githubCommitResponses, nil
}

func sleep(resp *http.Response) {
	log.Printf("[INFO]:\tGitHub responded with a status code: %d, server sleeping off for 1 hour...", resp.StatusCode)
	remaining, err := strconv.ParseInt(resp.Header.Get("x-ratelimit-reset"), 10, 64)
	if err != nil {
		time.Sleep(70 * time.Minute)
	} else {
		time.Sleep(time.Unix(remaining, 0).Sub(time.Now()))
	}
}
