package services

import (
	"encoding/json"
	"fmt"
	"github.com/djfemz/savannahTechTask/api/appErrors"
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/mappers"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CommitMonitorService struct {
	*CommitService
}

func NewCommitMonitorService(commitRepository *CommitService) *CommitMonitorService {
	return &CommitMonitorService{
		commitRepository,
	}
}

func (commitMonitorService *CommitMonitorService) FetchCommitData(page uint64) (githubCommitResponses *[]dtos.GitHubCommitResponse, err error) {
	resp, err := getData(os.Getenv("GITHUB_API_COMMIT_URL"), page, nil)
	if err != nil {
		log.Println("error sending request: ", err)
		return nil, err
	}
	githubCommitResponses, err = extractDataInto(resp, githubCommitResponses)
	return githubCommitResponses, nil
}

func getData(url string, page uint64, start *time.Time) (resp *http.Response, err error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	addHeadersToRequest(req, start, page)

	log.Println("req ", req)
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	log.Println("resp: ", resp)

	return resp, err
}

func (commitMonitorService *CommitMonitorService) StartJob() {
	job := cron.New()
	_, err := job.AddFunc("* * */1 * *", func() {
		log.Println("task in commit monitor service")
		go commitMonitorService.fetch(0)

	})
	if err != nil {
		log.Println("Error creating job: ", err)
	}
	log.Println("Starting new task...")
	job.Start()
}

func (commitMonitorService *CommitMonitorService) fetch(counter int) {
	data, err := commitMonitorService.FetchCommitData(uint64(counter))
	if err != nil {
		log.Println("Error fetching commits: ", err)
	}
	log.Println("data: ", data)
	if data != nil && len(*data) > 0 {
		commits := mappers.MapToCommits(data)
		err = commitMonitorService.repository.SaveAll(commits)
	}
}

func addHeadersToRequest(req *http.Request, start *time.Time, page uint64) {
	query := req.URL.Query()
	if page > 0 {
		query.Add("page", strconv.FormatUint(page, 10))
		query.Add("per_page", strconv.FormatUint(100, 10))
	}

	if start != nil {
		query.Add("since", start.String())
	}
	query.Add("Accept", "application/vnd.github+json")
	query.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AUTH_TOKEN")))
	req.URL.RawQuery = query.Encode()
	log.Println("req: ", req)
}

func extractDataInto[t any](resp *http.Response, into *t) (*t, error) {
	if resp.StatusCode != 200 {
		return nil, appErrors.NewCommitNotFoundError()
	}
	err := json.NewDecoder(resp.Body).Decode(&into)
	if err != nil {
		log.Println("Error reading response: ", err)
		return nil, err
	}
	return into, nil
}
