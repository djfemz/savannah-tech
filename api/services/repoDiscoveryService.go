package services

import (
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/models"
	"log"
	"os"
)

type RepoDiscoveryService struct {
	*GithubRepositoryService
}

func NewRepoDiscoveryService(service *GithubRepositoryService) *RepoDiscoveryService {
	return &RepoDiscoveryService{service}
}

func (repoDiscoveryService *RepoDiscoveryService) FetchRepoMetaData(errorChannel chan<- any) (githubRepository *dtos.GithubRepositoryResponse, err error) {
	resp, err := getData(os.Getenv("GITHUB_API_REPOSITORY_URL"), 0, nil)
	if err != nil {
		log.Println("Error fetching repository data: ", err)
		errorChannel <- err
		return nil, err
	}
	githubRepository, err = extractDataInto(resp, githubRepository)
	if err != nil {
		log.Println("Error extracting repository data from response: ", err)
		return nil, err
	}
	return githubRepository, err
}

func (repoDiscoveryService *RepoDiscoveryService) StartJob(doneChannel chan<- bool, errorChannel chan<- any) {
	go repoDiscoveryService.getRepoMetaData(doneChannel, errorChannel)
	log.Println("Starting new task...")
}

func (repoDiscoveryService *RepoDiscoveryService) getRepoMetaData(doneChannel chan<- bool, errorChannel chan<- any) {
	githubRepository, err := repoDiscoveryService.FetchRepoMetaData(errorChannel)
	if err != nil {
		errorChannel <- err
		log.Println("error fetching repo metadata: ", err)
		return
	}
	auxiliaryRepository := models.NewGithubRepository(githubRepository)
	if ok, _ := repoDiscoveryService.ExistsByName(githubRepository.Name); ok {
		auxiliaryRepository, _ = repoDiscoveryService.UpdateByName(githubRepository.Name, auxiliaryRepository)
		doneChannel <- true
		return
	}
	_, err = repoDiscoveryService.Save(auxiliaryRepository)
	if err != nil {
		log.Println("Error saving repository: ", err)
		errorChannel <- err
		return
	}
	doneChannel <- true
}
