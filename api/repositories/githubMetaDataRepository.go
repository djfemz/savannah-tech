package repositories

import (
	"github.com/djfemz/savannahTechTask/api/models"
	"gorm.io/gorm"
	"log"
)

type GithubAuxiliaryRepository interface {
	Save(repository *models.GithubRepository) (*models.GithubRepository, error)
	FindById(id uint) (*models.GithubRepository, error)
	FindByName(name string) (*models.GithubRepository, error)
	ExistsByName(name string) (bool, error)
	UpdateByName(name string, repository *models.GithubRepository) (*models.GithubRepository, error)
}

type GithubMetaDataRepository struct {
	*gorm.DB
}

func NewGithubAuxiliaryRepository(db *gorm.DB) GithubAuxiliaryRepository {
	return &GithubMetaDataRepository{db}
}

func (githubAuxRepo *GithubMetaDataRepository) Save(repository *models.GithubRepository) (*models.GithubRepository, error) {
	if err := githubAuxRepo.Create(repository).Error; err != nil {
		return nil, err
	}
	if err := githubAuxRepo.Last(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}

func (githubAuxRepo *GithubMetaDataRepository) FindById(id uint) (*models.GithubRepository, error) {
	repository := &models.GithubRepository{}
	if err := githubAuxRepo.Where("id=?", id).First(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}

func (githubAuxRepo *GithubMetaDataRepository) FindByName(name string) (*models.GithubRepository, error) {
	repository := &models.GithubRepository{}
	if err := githubAuxRepo.Where("name=?", name).First(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}

func (githubAuxRepo *GithubMetaDataRepository) ExistsByName(name string) (bool, error) {
	var repository *models.GithubRepository
	if err := githubAuxRepo.Where(&models.GithubRepository{Name: name}).First(&repository).Error; err != nil {

		log.Println("[ERROR:]\t ", err)
		return false, err
	} else if repository != nil && repository.Name == name {
		return true, nil
	}
	return false, nil
}

func (githubAuxRepo *GithubMetaDataRepository) UpdateByName(name string, repository *models.GithubRepository) (repo *models.GithubRepository, err error) {
	if err = githubAuxRepo.Model(&models.GithubRepository{}).Where("name = ?", name).Updates(repository).Error; err != nil {
		return nil, err
	}
	return repository, nil
}
