package repositories

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type CommitRepository struct {
	*gorm.DB
}

func NewCommitRepository(db *gorm.DB) *CommitRepository {
	return &CommitRepository{db}
}

func ConnectToDatabase() (*gorm.DB, error) {
	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 64)
	if err != nil {
		log.Fatal("Error reading port: ", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d TimeZone=Africa/Lagos sslmode=disable", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}
	return db, nil
}
