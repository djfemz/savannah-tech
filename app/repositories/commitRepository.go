package repositories

import (
	"github.com/djfemz/savannahTechTask/app/models"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
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
}

type AppCommitRepository struct {
	*gorm.DB
}

func NewCommitRepository(db *gorm.DB) CommitRepository {
	return &AppCommitRepository{db}
}

func (appCommitRepository *AppCommitRepository) Save(commit *models.Commit) (*models.Commit, error) {
	var author *models.Author
	appCommitRepository.Where(&models.Author{Username: commit.Author.Username}).First(author)

	if err := appCommitRepository.DB.Create(commit).Error; err != nil {
		return nil, err
	}
	if err := appCommitRepository.Last(commit).Error; err != nil {
		return nil, err
	}
	return commit, nil
}

func (appCommitRepository *AppCommitRepository) FindById(id uint) (*models.Commit, error) {
	foundCommit := &models.Commit{}
	if err := appCommitRepository.Where("id=?", id).First(foundCommit).Error; err != nil {
		return nil, err
	}
	return foundCommit, nil
}

func (appCommitRepository *AppCommitRepository) FindAllByDateSince(since *time.Time) (commits []*models.Commit, err error) {
	if err := appCommitRepository.Where("date BETWEEN ? AND ?", since, time.Now()).Find(&commits).Error; err != nil {
		return nil, err
	}
	return
}

func (appCommitRepository *AppCommitRepository) FindMostRecentCommit() (commit *models.Commit, err error) {
	if err = appCommitRepository.Order("date desc").First(commit).Error; err != nil {
		return nil, err
	}
	return
}

func (appCommitRepository *AppCommitRepository) FindAll() (commits []*models.Commit, err error) {
	if err := appCommitRepository.Find(&commits).Error; err != nil {
		log.Println("error finding commits: ", err)
		return nil, err
	}
	return commits, nil
}

func (appCommitRepository *AppCommitRepository) SaveAll(commits []*models.Commit) error {
	if err := appCommitRepository.DB.Save(commits).Error; err != nil {
		return err
	}
	return nil
}

func (appCommitRepository *AppCommitRepository) FindTopCommitAuthors(size int) ([]*models.Author, error) {
	var authors []*models.Author
	limit := strconv.Itoa(size)
	sql := "SELECT email, username, count(*) as commits from authors GROUP BY email, username ORDER BY count(*) DESC LIMIT " + limit
	log.Println("sql: "+sql, size)
	if err := appCommitRepository.Raw(sql).Limit(size).Find(&authors).Error; err != nil {
		return nil, err
	}
	log.Println("authors: ", authors)
	return authors, nil
}

func (appCommitRepository *AppCommitRepository) FindCommitsForRepoByName(name string) ([]*models.Commit, error) {
	var commits []*models.Commit
	if err := appCommitRepository.Find(&models.Commit{RepoName: name}).First(commits).Error; err != nil {
		return nil, err
	}
	return commits, nil
}
