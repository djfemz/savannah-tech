package services

import (
	"encoding/json"
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/models"
	"github.com/djfemz/savannahTechTask/app/repositories"
	"log"
	"net/http"
	"os"
)

type CommitService struct {
	repository *repositories.CommitRepository
}

func NewCommitService(repository *repositories.CommitRepository) *CommitService {
	return &CommitService{
		repository: repository,
	}
}

func (commitService *CommitService) GetAllCommits() (responses []*dtos.CommitResponse, err error) {
	var commits []*models.Commit
	if err = commitService.repository.Find(&commits).Error; err != nil {
		log.Println("error fetching commits: ", err)
	}
	responses = mapToCommitResponses(commits)
	return responses, err
}

func FetchCommits() {
	db, err := repositories.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	commitRepository := repositories.NewCommitRepository(db)
	commits, err := getCommits()
	usersCommits := mapToCommits(commits)
	log.Println("users commits: ", usersCommits)
	if err = commitRepository.Create(&usersCommits).Error; err != nil {
		log.Println("Error adding commits to database: ", err)
		return
	}
}

func getCommits() (commits []*dtos.GitHubCommitResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
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
