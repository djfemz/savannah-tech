package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/djfemz/savannahTechTask/api/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/mappers"
)

type CommitMonitorService struct {
	*CommitManager
}

func NewCommitMonitorService(commitManager *CommitManager) *CommitMonitorService {
	return &CommitMonitorService{
		commitManager,
	}
}

func (commitMonitorService *CommitMonitorService) FetchCommitData() (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
	commit, err := commitMonitorService.GetMostRecentCommit()
	if err != nil {
		log.Printf("[Error: %v]", err)
	}
	githubCommitResponses, err = commitMonitorService.fetchAllCommits(githubCommitResponses, &commit.Date)
	return githubCommitResponses, err
}

func getData(url string, page int, start *time.Time) (resp *http.Response, err error) {
	client := http.Client{}
	url = addHeadersToRequest(url, page, start)
	log.Println("url: ", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AUTH_TOKEN")))
	req.Header.Add("Accept", utils.ACCEPT_HEADER_VALUE)
	if err != nil {
		return nil, err
	}
	log.Println("req: ", req)
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	// TODO: queue up the failed requests for retries
	return resp, err
}

func (commitMonitorService *CommitMonitorService) StartJob() {
	log.Println("Repo monitoring started...")
	commitMonitorService.fetch()
}

func (commitMonitorService *CommitMonitorService) fetch() {
	repoName := os.Getenv("REPO_NAME")
	repository, _ := commitMonitorService.FindByName(repoName)
	data, err := commitMonitorService.FetchCommitData()
	if err != nil {
		log.Println("Error fetching commits: ", err)
	}

	if data != nil && len(*data) > 0 {
		commits := mappers.MapToCommits(data, repository)
		err = commitMonitorService.repository.SaveAll(commits)
		if err != nil {
			log.Println("error saving commits: ", err)
		}
	}
}

func addHeadersToRequest(url string, page int, start *time.Time) string {
	if page > 0 {
		url = fmt.Sprintf("%s%s%s%s%s", url, "?page=", strconv.Itoa(page),
			"&per_page=", strconv.Itoa(MAX_RECORDS_PER_PAGE))
	}

	if start != nil {
		url = fmt.Sprintf("%s&since=%s&until=%s", url, start.String(), time.Now().String())
	}
	return url
}

func extractDataInto[t any](resp *http.Response, into *t) (*t, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request responded to with status: %d", resp.StatusCode))
	}
	err := json.NewDecoder(resp.Body).Decode(&into)
	if err != nil {
		log.Println("Error reading response: ", err)
		return nil, err
	}
	return into, nil
}
