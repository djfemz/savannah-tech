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

// TODO: work on calling fetch operation in endpoint that adds repoName
func getData(url string, page int, start *time.Time) (resp *http.Response, err error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	log.Println("req: ", req)
	addHeadersTo(req)
	addParamsTo(req, page, start)
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func addParamsTo(req *http.Request, page int, start *time.Time) {
	query := req.URL.Query()
	isPageNumberValid := page >= MINIMUM_ALLOWED_PAGE_NUMBER
	if isPageNumberValid {
		query.Add("page", strconv.Itoa(page))
		query.Add("per_page", strconv.Itoa(MAX_RECORDS_PER_PAGE))
	}
	isStartTimeValid := start != nil
	if isStartTimeValid {
		query.Add("since", start.String())
		query.Add("until", time.Now().String())
	}
	req.URL.RawQuery = query.Encode()
}

func addHeadersTo(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AUTH_TOKEN")))
	req.Header.Add("Accept", utils.ACCEPT_HEADER_VALUE)
}

func (commitMonitorService *CommitMonitorService) StartJob() {
	log.Println("[INFO:]\tRepo monitoring started...")
	commitMonitorService.fetch()
}

func (commitMonitorService *CommitMonitorService) fetch() {
	repoName := os.Getenv("REPO_NAME")
	repository, _ := commitMonitorService.FindByName(repoName)
	data, err := commitMonitorService.FetchCommitData()
	if err != nil {
		log.Println("[ERROR:]\tError fetching commits: ", err)
	}

	if data != nil && len(*data) > 0 {
		commits := mappers.MapToCommits(data, repository)
		err = commitMonitorService.repository.SaveAll(commits)
		if err != nil {
			log.Println("[ERROR:]\terror saving commits: ", err)
		}
	}
}

func extractDataInto[t any](resp *http.Response, into *t) (*t, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request responded to with status: %d", resp.StatusCode))
	}
	err := json.NewDecoder(resp.Body).Decode(&into)
	if err != nil {
		log.Println("[ERROR:]\tError reading response: ", err)
		return nil, err
	}
	return into, nil
}
