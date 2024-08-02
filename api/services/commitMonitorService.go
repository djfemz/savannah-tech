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
	if page > 0 {
		url = fmt.Sprintf("%s%s%s%s%s", url, "?page=", strconv.Itoa(page),
			"&per_page=", strconv.Itoa(MAX_RECORDS_PER_PAGE))
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	addHeadersToRequest(req, start)
	log.Println("req: ", req)
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	log.Println("resp: ", resp)
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

func addHeadersToRequest(req *http.Request, start *time.Time) {
	query := req.URL.Query()
	if start != nil {
		query.Add("since", start.String())
		query.Add("until", time.Now().String())
	}
	query.Add("Accept", utils.ACCEPT_HEADER_VALUE)
	query.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AUTH_TOKEN")))
	req.URL.RawQuery = query.Encode()
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
