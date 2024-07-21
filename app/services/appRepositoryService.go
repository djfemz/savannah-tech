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

type AppRepositoryService struct {
	*repositories.AppRepoRepository
}

func (appRepositoryService *AppRepositoryService) FetchRepositoryMetaData() {
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
	appRepository := models.NewAppRepository(repository)
	if err = appRepositoryService.AppRepoRepository.Create(appRepository).Error; err != nil {
		log.Println("Error saving repository: ", err)
	}
}

func (appRepositoryService *AppRepositoryService) GetRepository(name string) (repository *dtos.RepositoryResponse, err error) {
	appRepository := &models.AppRepository{}
	err = appRepositoryService.AppRepoRepository.Where(&models.AppRepository{Name: name}).Find(appRepository).Error
	if err != nil {
		return nil, err
	}
	return models.NewRepositoryResponse(appRepository), nil
}

func NewGithubRepoService(appRepoRepository *repositories.AppRepoRepository) *AppRepositoryService {
	return &AppRepositoryService{appRepoRepository}
}
