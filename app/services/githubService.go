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

func (githubService *GithubService) FetchCommits() (githubUserCommits []*dtos.GitHubCommitResponse, err error) {
	var commits []*models.Commit
	var commit = &models.Commit{}
	commit, _ = githubService.GetMostRecentCommit()
	githubUserCommits, err = getCommits(commit)
	if err != nil {
		return nil, err
	}
	commits = mapToCommits(githubUserCommits)
	if err = githubService.repository.SaveAll(&commits); err != nil {
		log.Println("Error adding githubUserCommits to database: ", err)
		return nil, err
	}
	return githubUserCommits, nil
}

// TODO: remove hardcoded env variables
func (githubService *GithubService) FetchRepositoryMetaData() {
	repository := dtos.NewGithubRepositoryResponse()
	req, err := http.NewRequest(http.MethodGet, os.Getenv("GITHUB_API_REPOSITORY_URL"), nil)
	if err != nil {
		log.Println("Error creating request: ", err)
	}
	client := &http.Client{}
	jsonResponse, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	err = json.NewDecoder(jsonResponse.Body).Decode(repository)
	if err != nil {
		log.Fatal("Error reading response: ", err)
	}
	appRepository := models.NewGithubAuxiliaryRepository(repository)
	repoName := os.Getenv("REPO_NAME")
	isExistingRepo, _ := githubService.ExistsByName(repoName)
	if isExistingRepo {
		repo, _ := githubService.UpdateByName(repoName, appRepository)
		log.Println("repository updated: ", repo)
		return
	}
	_, err = githubService.GithubAuxiliaryRepository.Save(appRepository)
	if err != nil {
		log.Println("Error saving repository: ", err)
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
