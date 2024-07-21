package repositories

import (
	"fmt"
	"github.com/djfemz/savannahTechTask/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

type CommitRepository interface {
	Save(commit *models.Commit) (*models.Commit, error)
	FindById(id uint) (*models.Commit, error)
	FindAllByDateSince(since *time.Time) ([]*models.Commit, error)
	FindAll() ([]*models.Commit, error)
	SaveAll(commits *[]*models.Commit) error
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

func (appCommitRepository *AppCommitRepository) FindAll() (commits []*models.Commit, err error) {
	if err := appCommitRepository.Find(&commits).Error; err != nil {
		log.Println("error finding commits: ", err)
		return nil, err
	}
	return commits, nil
}

func (appCommitRepository *AppCommitRepository) SaveAll(commits *[]*models.Commit) error {
	if err := appCommitRepository.Create(commits).Error; err != nil {
		return err
	}
	return nil
}

func ConnectToDatabase() (*gorm.DB, error) {
	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 64)
	if err != nil {
		log.Fatal("Error reading port: ", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Africa/Lagos sslmode=disable", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	err = db.AutoMigrate(&models.Commit{}, &models.Author{}, &models.GithubAuxiliaryRepository{})
	if err != nil {
		log.Fatal("error migrating: ", err)
	}
	return db, nil
}
