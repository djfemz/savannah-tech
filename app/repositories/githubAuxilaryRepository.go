package repositories

import (
	"github.com/djfemz/savannahTechTask/app/models"
	"gorm.io/gorm"
)

type GithubAuxiliaryRepository interface {
	Save(repository *models.GithubAuxiliaryRepository) (*models.GithubAuxiliaryRepository, error)
	FindById(id uint) (*models.GithubAuxiliaryRepository, error)
	FindByName(name string) (*models.GithubAuxiliaryRepository, error)
}

type AppGithubAuxiliaryRepository struct {
	*gorm.DB
}

func NewGithubAuxiliaryRepository(db *gorm.DB) GithubAuxiliaryRepository {
	return &AppGithubAuxiliaryRepository{db}
}

func (githubAuxRepo *AppGithubAuxiliaryRepository) Save(repository *models.GithubAuxiliaryRepository) (*models.GithubAuxiliaryRepository, error) {
	if err := githubAuxRepo.Create(repository).Error; err != nil {
		return nil, err
	}
	if err := githubAuxRepo.Last(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}

func (githubAuxRepo *AppGithubAuxiliaryRepository) FindById(id uint) (*models.GithubAuxiliaryRepository, error) {
	repository := &models.GithubAuxiliaryRepository{}
	if err := githubAuxRepo.Where("id=?", id).First(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}

func (githubAuxRepo *AppGithubAuxiliaryRepository) FindByName(name string) (*models.GithubAuxiliaryRepository, error) {
	repository := &models.GithubAuxiliaryRepository{}
	if err := githubAuxRepo.Where("name=?", name).First(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}
