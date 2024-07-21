package test

import (
	"github.com/djfemz/savannahTechTask/app/mocks"
	"github.com/djfemz/savannahTechTask/app/models"
	"github.com/djfemz/savannahTechTask/app/repositories"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func TestDatabaseConnection(t *testing.T) {
	db, err := repositories.ConnectToDatabase()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestGetLastCommit(t *testing.T) {
	db, err := repositories.ConnectToDatabase()
	assert.Nil(t, err)
	var commit = &models.Commit{}
	db.Order("date desc").First(commit)
	log.Println("commit: ", *commit)
	assert.NotNil(t, commit)
}

func TestFindLastCommitByDate(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindMostRecentCommit").Return(testCommits[0], nil)
	commit, err := commitRepository.FindMostRecentCommit()
	assert.Nil(t, err)
	assert.NotNil(t, commit)
}

func TestFindTopCommitAuthors(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("FindMostRecentCommit").Return(testCommits[0], nil)
}

func TestAddSameCommitFails(t *testing.T) {
	//db, err := repositories.ConnectToDatabase()
	//assert.Nil(t, err)
	//err = db.Create(&models.Commit{CommitHash: "12345abc"}).Error
	//assert.Nil(t, err)
	//err = db.Create(&models.Commit{CommitHash: "12345abc"}).Error
	//log.Println("Error: ", err)
	//assert.NotNil(t, err)

}

func TestAddBatchCommit(t *testing.T) {
	//db, err := repositories.ConnectToDatabase()
	//assert.Nil(t, err)
	//commits := []*models.Commit{{CommitHash: "12345abp"}, {CommitHash: "12345efg"}, {CommitHash: "12345xyz"}}
	//err = db.Create(commits).Error
	//log.Println("Error: ", err)
	//assert.Nil(t, err)

}
