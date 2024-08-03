package repositories

import (
	"log"
	"strconv"
	"time"

	"github.com/djfemz/savannahTechTask/api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommitRepository interface {
	Save(commit *models.Commit) (*models.Commit, error)
	FindById(id uint) (*models.Commit, error)
	FindAllByDateSince(since *time.Time) ([]*models.Commit, error)
	FindAll() ([]*models.Commit, error)
	SaveAll(commits []*models.Commit) error
	FindMostRecentCommit() (*models.Commit, error)
	FindTopCommitAuthors(size int) ([]*models.Author, error)
	FindCommitsForRepoByName(name string) ([]*models.Commit, error)
	CountCommits() (int64, error)
}

type AppCommitRepository struct {
	*gorm.DB
}

func NewCommitRepository(db *gorm.DB) CommitRepository {
	return &AppCommitRepository{db}
}

func (appCommitRepository *AppCommitRepository) Save(commit *models.Commit) (*models.Commit, error) {

	if err := appCommitRepository.Create(commit).Error; err != nil {
		return nil, err
	}
	if err := appCommitRepository.Last(commit).Error; err != nil {
		return nil, err
	}
	return commit, nil
}

func (appCommitRepository *AppCommitRepository) FindById(id uint) (*models.Commit, error) {
	foundCommit := &models.Commit{}
	if err := appCommitRepository.Preload(clause.Associations).Where("id=?", id).First(foundCommit).Error; err != nil {
		return nil, err
	}
	return foundCommit, nil
}

func (appCommitRepository *AppCommitRepository) FindAllByDateSince(since *time.Time) (commits []*models.Commit, err error) {
	if err := appCommitRepository.Preload(clause.Associations).Where("created_at BETWEEN ? AND ?", since, time.Now()).Find(&commits).Error; err != nil {
		return nil, err
	}
	return
}

func (appCommitRepository *AppCommitRepository) FindMostRecentCommit() (*models.Commit, error) {
	commit := models.Commit{}
	if err := appCommitRepository.Preload(clause.Associations).Order("committed_at desc").First(&commit).Error; err != nil {
		log.Println("[ERROR:]\t error finding most recent commit ", err)
		return nil, err
	}
	return &commit, nil
}

func (appCommitRepository *AppCommitRepository) FindAll() (commits []*models.Commit, err error) {
	if err := appCommitRepository.Preload(clause.Associations).Find(&commits).Error; err != nil {
		return nil, err
	}
	return commits, nil
}

func (appCommitRepository *AppCommitRepository) SaveAll(commits []*models.Commit) error {
	log.Println("[INFO:]\t saving data")
	if err := appCommitRepository.DB.Save(commits).Error; err != nil {
		log.Println("[ERROR:]\t error saving commits", err)
		return err
	}
	return nil
}

func (appCommitRepository *AppCommitRepository) FindTopCommitAuthors(size int) ([]*models.Author, error) {
	var authors []*models.Author
	limit := strconv.Itoa(size)
	sql := "SELECT email, username, count(*) as commits from authors GROUP BY email, username ORDER BY count(*) DESC LIMIT " + limit
	if err := appCommitRepository.Raw(sql).Limit(size).Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (appCommitRepository *AppCommitRepository) FindCommitsForRepoByName(name string) ([]*models.Commit, error) {
	var commits []*models.Commit
	if err := appCommitRepository.Preload(clause.Associations).Where(&models.Commit{RepoName: name}).Find(&commits).Error; err != nil {
		return nil, err
	}
	return commits, nil
}

func (appCommitRepository *AppCommitRepository) CountCommits() (int64, error) {
	var commits []models.Commit
	var count int64
	err := appCommitRepository.Model(&commits).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
