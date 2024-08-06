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

var repoUrl string

func NewRepoDiscoveryService(service *GithubRepositoryService) *RepoDiscoveryService {
	repoUrl = os.Getenv("GITHUB_API_REPOSITORY_URL")
	return &RepoDiscoveryService{service}
}

func (repoDiscoveryService *RepoDiscoveryService) FetchRepoMetaData(errorChannel chan<- any) (githubRepository *dtos.GithubRepositoryResponse, err error) {
	resp, err := getData(repoUrl, 0, nil)
	if err != nil {
		log.Println("[Error:]\t", err)
		errorChannel <- err
		return nil, err
	}
	githubRepository, err = extractDataInto(resp, githubRepository)
	if err != nil {
		log.Println("[Error:]\textracting repository data from response: ", err)
		return nil, err
	}
	return githubRepository, err
}

func (repoDiscoveryService *RepoDiscoveryService) StartJob(doneChannel chan<- bool, errorChannel chan<- any) {
	log.Println("[INFO:]\t Starting task to fetch Repository Metadata..")
	go repoDiscoveryService.getRepoMetaData(doneChannel, errorChannel)
}

func (repoDiscoveryService *RepoDiscoveryService) getRepoMetaData(doneChannel chan<- bool, errorChannel chan<- any) {
	githubRepository, err := repoDiscoveryService.FetchRepoMetaData(errorChannel)
	if err != nil {
		errorChannel <- err
		log.Println("[ERROR:]\terror fetching repo metadata: ", err)
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
		log.Println("[ERROR:]\tError saving repository: ", err)
		errorChannel <- err
		return
	}
	doneChannel <- true
}
