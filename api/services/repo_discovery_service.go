package services

import (
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/models"
	"log"
	"os"
)

type RepoDiscoveryService struct {
	*GithubRepositoryService
	*log.Logger
}

var repoUrl string

func NewRepoDiscoveryService(service *GithubRepositoryService, logger *log.Logger) *RepoDiscoveryService {
	repoUrl = os.Getenv("GITHUB_API_BASE_URL")
	return &RepoDiscoveryService{service, logger}
}

func (repoDiscoveryService *RepoDiscoveryService) FetchRepoMetaData(errorChannel chan<- any) (githubRepository *dtos.GithubRepositoryResponse, err error) {
	resp, err := getData(repoUrl, 0, nil)
	if err != nil {
		repoDiscoveryService.Println("\t", err)
		errorChannel <- err
		return nil, err
	}
	githubRepository, err = extractDataInto(resp, githubRepository)
	if err != nil {
		repoDiscoveryService.Println("\textracting repository data from response: ", err)
		return nil, err
	}
	return githubRepository, err
}

func (repoDiscoveryService *RepoDiscoveryService) StartJob(doneChannel chan<- bool, errorChannel chan<- any) {
	repoDiscoveryService.Println("\t Starting task to pullCommitDataFromGithub Repository Metadata..")
	go repoDiscoveryService.getRepoMetaData(doneChannel, errorChannel)
}

func (repoDiscoveryService *RepoDiscoveryService) getRepoMetaData(doneChannel chan<- bool, errorChannel chan<- any) {
	githubRepository, err := repoDiscoveryService.FetchRepoMetaData(errorChannel)
	if err != nil {
		errorChannel <- err
		repoDiscoveryService.Println("\terror fetching repo metadata: ", err)
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
		repoDiscoveryService.Println("\tError saving repository: ", err)
		errorChannel <- err
		return
	}
	doneChannel <- true
}
