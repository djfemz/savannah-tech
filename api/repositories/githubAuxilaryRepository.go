package repositories

import (
	"github.com/djfemz/savannahTechTask/api/models"
	"gorm.io/gorm"
)

type GithubAuxiliaryRepository interface {
	Save(repository *models.GithubAuxiliaryRepository) (*models.GithubAuxiliaryRepository, error)
	FindById(id uint) (*models.GithubAuxiliaryRepository, error)
	FindByName(name string) (*models.GithubAuxiliaryRepository, error)
	ExistsByName(name string) (bool, error)
	UpdateByName(name string, repository *models.GithubAuxiliaryRepository) (*models.GithubAuxiliaryRepository, error)
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

func (githubAuxRepo *AppGithubAuxiliaryRepository) ExistsByName(name string) (bool, error) {
	var repository *models.GithubAuxiliaryRepository
	if err := githubAuxRepo.Where(&models.GithubAuxiliaryRepository{Name: name}).First(repository).Error; err != nil {
		return false, err
	} else if repository.Name == name {
		return true, nil
	}
	return false, nil
}

func (githubAuxRepo *AppGithubAuxiliaryRepository) UpdateByName(name string, repository *models.GithubAuxiliaryRepository) (repo *models.GithubAuxiliaryRepository, err error) {
	if err = githubAuxRepo.Model(&models.GithubAuxiliaryRepository{}).Where("name = ?", name).Updates(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}
