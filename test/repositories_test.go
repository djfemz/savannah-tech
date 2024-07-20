package test

import (
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
