package services

import (
	"encoding/json"
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/models"
	"github.com/djfemz/savannahTechTask/app/repositories"
	"log"
	"net/http"
	"os"
	"time"
)

type CommitService struct {
	repository *repositories.CommitRepository
}

func NewCommitService(repository *repositories.CommitRepository) *CommitService {
	return &CommitService{
		repository: repository,
	}
}

// TODO: Add Caching Strategies
func (commitService *CommitService) GetAllCommits() (responses []*dtos.CommitResponse, err error) {
	var commits []*models.Commit
	if err = commitService.repository.Find(&commits).Error; err != nil {
		log.Println("error fetching commits: ", err)
	}
	responses = mapToCommitResponses(commits)
	return responses, err
}

func (commitService *CommitService) GetCommitsByDateSince(since time.Time) (response []*dtos.CommitResponse, err error) {
	commits := make([]*models.Commit, 0)
	err = commitService.repository.Where("date BETWEEN ? AND ?", since, time.Now()).Find(&commits).Error
	if err != nil {
		return nil, err
	}
	return mapToCommitResponses(commits), err
}

func FetchCommits() {
	var commits []*models.Commit
	var githubUserCommits []*dtos.GitHubCommitResponse
	var commit = &models.Commit{}
	var err error
	db, err := repositories.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	db.Order("date desc").First(commit)
	githubUserCommits, err = getCommits(commit)
	commits = mapToCommits(githubUserCommits)
	commitRepository := repositories.NewCommitRepository(db)
	if err = commitRepository.Create(&commits).Error; err != nil {
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
		since := commit.Date.Add(1*time.Minute).Add(1*time.Nanosecond).Format(time.RFC3339) + "Z"
		query := req.URL.Query()
		query.Add("since", since)
		req.URL.RawQuery = query.Encode()
	}
	client := &http.Client{}
	jsonResponse, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	err = json.NewDecoder(jsonResponse.Body).Decode(&commits)
	if err != nil {
		log.Fatal("Error reading response: ", err)
	}
	return commits, err
}

func mapToCommits(commits []*dtos.GitHubCommitResponse) []*models.Commit {
	var usersCommits = make([]*models.Commit, 0)
	for _, commit := range commits {
		userCommit := models.NewCommitFromGithubCommitResponse(commit)
		usersCommits = append(usersCommits, userCommit)
	}
	return usersCommits
}

func mapToCommitResponses(commits []*models.Commit) []*dtos.CommitResponse {
	var usersCommits = make([]*dtos.CommitResponse, 0)
	for _, commit := range commits {
		userCommit := models.NewCommitResponse(commit)
		usersCommits = append(usersCommits, userCommit)
	}
	return usersCommits
}
