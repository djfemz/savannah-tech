package services

import (
	"encoding/json"
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/models"
	"log"
	"net/http"
	"os"
	"time"
)

type GithubService struct {
	*CommitService
	*GithubRepositoryService
}

func NewGithubService(commitService *CommitService, repositoryService *GithubRepositoryService) *GithubService {
	return &GithubService{
		commitService,
		repositoryService,
	}
}

func (githubService *GithubService) FetchCommits() {
	var commits []*models.Commit
	var githubUserCommits []*dtos.GitHubCommitResponse
	var commit = &models.Commit{}
	var err error
	if err != nil {
		log.Println("Error connecting to database: ", err)
	}
	//TODO: implement a get latest commit --> db.Order("date desc").First(commit)
	githubUserCommits, err = getCommits(commit)
	commits = mapToCommits(githubUserCommits)

	if err = githubService.repository.SaveAll(&commits); err != nil {
		log.Println("Error adding githubUserCommits to database: ", err)
		return
	}
}

func getCommits(commit *models.Commit) (commits []*dtos.GitHubCommitResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), nil)
	if err != nil {
		log.Println("Error creating request: ", err)
	}
	if commit != nil {
		addQueryParamToRequest(commit, req)
	}
	client := &http.Client{}
	jsonResponse, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return
	}
	err = json.NewDecoder(jsonResponse.Body).Decode(&commits)
	if err != nil {
		log.Println("Error reading response: ", err)
		return
	}
	return commits, err
}

func addQueryParamToRequest(commit *models.Commit, req *http.Request) {
	since := commit.Date.Add(1*time.Minute).Add(1*time.Nanosecond).Format(time.RFC3339) + "Z"
	query := req.URL.Query()
	query.Add("since", since)
	req.URL.RawQuery = query.Encode()
}
