package services

import (
	"github.com/djfemz/savannahTechTask/api/appErrors"
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/models"
	"github.com/djfemz/savannahTechTask/api/repositories"
)

type GithubRepositoryService struct {
	repositories.GithubAuxiliaryRepository
}

func NewGithubRepoMetadataService(appRepoRepository repositories.GithubAuxiliaryRepository) *GithubRepositoryService {
	return &GithubRepositoryService{appRepoRepository}
}

func (appRepositoryService *GithubRepositoryService) GetRepositoryBy(name string) (repository *dtos.RepositoryResponse, err error) {
	appRepository, err := appRepositoryService.GithubAuxiliaryRepository.FindByName(name)
	if err != nil {
		return nil, appErrors.NewRepositoryNotFoundError()
	}
	return models.NewRepositoryResponse(appRepository), nil
}
