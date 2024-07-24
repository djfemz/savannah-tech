package services

import (
	"log"
	"os"
	"time"

	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/models"
	"github.com/robfig/cron/v3"
)

type RepoDiscoveryService struct {
	*GithubRepositoryService
}

func NewRepoDiscoveryService(service *GithubRepositoryService) *RepoDiscoveryService {
	return &RepoDiscoveryService{service}
}

func (repoDiscoveryService *RepoDiscoveryService) FetchRepoMetaData() (githubRepository *dtos.GithubRepositoryResponse, err error) {
	resp, err := getData(os.Getenv("GITHUB_API_REPOSITORY_URL"), 0, nil)
	if err != nil {
		log.Println("Error fetching repository data: ", err)
		return nil, err
	}
	githubRepository, err = extractDataInto(resp, githubRepository)
	if err != nil {
		log.Println("Error extracting repository data from response: ", err)
		return nil, err
	}
	return githubRepository, err
}

func (repoDiscoveryService *RepoDiscoveryService) StartJob() {
	job := cron.New()
	_, err := job.AddFunc("* * */1 * *", func() {
		time.Sleep(3 * time.Second)
		go repoDiscoveryService.fetch()
	})
	if err != nil {
		log.Println("Error creating job: ", err)
	}
	log.Println("Starting new task...")
	job.Start()
}

func (repoDiscoveryService *RepoDiscoveryService) fetch() {
	githubRepository, err := repoDiscoveryService.FetchRepoMetaData()
	if err != nil {
		log.Println("error: ", err)
		return
	}
	auxiliaryRepository := models.NewGithubAuxiliaryRepository(githubRepository)
	if ok, _ := repoDiscoveryService.ExistsByName(githubRepository.Name); ok {
		auxiliaryRepository, _ = repoDiscoveryService.UpdateByName(githubRepository.Name, auxiliaryRepository)
	}
	auxiliaryRepository, _ = repoDiscoveryService.Save(auxiliaryRepository)
}
