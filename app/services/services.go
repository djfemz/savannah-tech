package services

import (
	"encoding/json"
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/repositories"
	"log"
	"net/http"
	"os"
)

const (
	API_KEY_VALUE          = "api-key"
	CONTENT_TYPE_KEY       = "Content-Type"
	APPLICATION_JSON_VALUE = "application/json"
	ACCEPT_HEADER_KEY      = "accept"
)

type CommitService struct {
	repository *repositories.CommitRepository
}

func NewCommitService(repository *repositories.CommitRepository) *CommitService {
	return &CommitService{
		repository: repository,
	}
}

func (commitService CommitService) GetAllCommits() ([]*dtos.CommitResponse, error) {
	return nil, nil
}

func FetchCommits() {
	db, err := repositories.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	commitRepository := repositories.NewCommitRepository(db)
	commits, err := getCommits()
	if err = commitRepository.Create(&commits).Error; err != nil {
		log.Println("Error adding commits to database: ", err)
		return
	}
}

func getCommits() ([]*dtos.GitHubCommitResponse, error) {
	var err error
	req, err := http.NewRequest(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}
	client := &http.Client{}
	jsonResponse, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	var commits []*dtos.GitHubCommitResponse
	err = json.NewDecoder(jsonResponse.Body).Decode(&commits)
	if err != nil {
		log.Fatal("Error reading response: ", err)
	}
	return commits, err
}

func addHeadersTo(req *http.Request) {
	req.Header.Add(API_KEY_VALUE, os.Getenv("MAIL_API_KEY"))
	req.Header.Add(CONTENT_TYPE_KEY, APPLICATION_JSON_VALUE)
	req.Header.Add(ACCEPT_HEADER_KEY, APPLICATION_JSON_VALUE)
}
