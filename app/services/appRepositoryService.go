package services

import (
	dtos "github.com/djfemz/savannahTechTask/app/dtos/responses"
	"github.com/djfemz/savannahTechTask/app/models"
	"github.com/djfemz/savannahTechTask/app/repositories"
)

type GithubRepositoryService struct {
	repositories.GithubAuxiliaryRepository
}

func NewGithubRepoService(appRepoRepository repositories.GithubAuxiliaryRepository) *GithubRepositoryService {
	return &GithubRepositoryService{appRepoRepository}
}

func (appRepositoryService *GithubRepositoryService) GetRepositoryBy(name string) (repository *dtos.RepositoryResponse, err error) {
	appRepository, err := appRepositoryService.GithubAuxiliaryRepository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return models.NewRepositoryResponse(appRepository), nil
}
