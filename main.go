package main

import (
	"github.com/djfemz/savannahTechTask/app/controllers"
	"github.com/djfemz/savannahTechTask/app/repositories"
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file: ", err)
	}
}

func main() {
	db, _ := repositories.ConnectToDatabase()
	commitController := controllers.NewCommitController(services.NewCommitService(repositories.NewCommitRepository(db)))
	// TODO: add channels to communicate
	go startJob()
	router := gin.Default()
	router.GET("/api/v1/commits", commitController.GetTopCommitAuthors)
	router.GET("/api/v1/commits/:repo", commitController.GetCommitsForRepository)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server on port: ", 8080)
	}
}

func startJob() {
	db, _ := repositories.ConnectToDatabase()
	githubService := services.NewGithubService(services.NewCommitService(repositories.NewCommitRepository(db)), services.NewGithubRepoService(repositories.NewGithubAuxiliaryRepository(db)))
	job := cron.New()
	_, err := job.AddFunc("* */1 * * *", func() {
		_, err := githubService.FetchCommits()
		if err != nil {
			log.Println("Error fetching commits: ", err)
		}
	})
	if err != nil {
		log.Println("Error creating job: ", err)
	}
	log.Println("Starting new task...")
	job.Start()
}
