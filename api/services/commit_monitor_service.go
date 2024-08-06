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

var authToken string

func init() {

}

type CommitMonitorService struct {
	*CommitManager
	*log.Logger
}

func NewCommitMonitorService(commitManager *CommitManager, logger *log.Logger) *CommitMonitorService {
	authToken = os.Getenv("AUTH_TOKEN")
	return &CommitMonitorService{
		commitManager,
		logger,
	}
}

func (commitMonitorService *CommitMonitorService) FetchCommitData() (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
	githubCommitResponses, err = commitMonitorService.fetchAllCommits(githubCommitResponses, nil)
	return githubCommitResponses, err
}

func getData(url string, page int64, start *time.Time) (resp *http.Response, err error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	addHeadersTo(req)
	addParamsTo(req, int(page), start)
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
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))
	req.Header.Add("Accept", utils.ACCEPT_HEADER_VALUE)
}

func (commitMonitorService *CommitMonitorService) StartJob() {
	commitMonitorService.Println("\tRepo monitoring started...")
	commitMonitorService.pullCommitDataFromGithub()
}

func (commitMonitorService *CommitMonitorService) pullCommitDataFromGithub() {
	repository, _ := commitMonitorService.FindByName(repoName)
	data, err := commitMonitorService.FetchCommitData()
	if err != nil {
		commitMonitorService.Println("\tError fetching commits: ", err)
	}

	if data != nil && len(*data) > 0 {
		commits := mappers.MapToCommits(data, repository)
		err = commitMonitorService.repository.SaveAll(commits)
		if err != nil {
			commitMonitorService.Println("\terror saving commits: ", err)
		}
	}
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
