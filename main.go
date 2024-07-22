package main

import (
	"github.com/djfemz/savannahTechTask/app/controllers"
	"github.com/djfemz/savannahTechTask/app/repositories"
	"github.com/djfemz/savannahTechTask/app/services"
	_ "github.com/djfemz/savannahTechTask/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"os"
)

func init() {
	err := godotenv.Load("example.env")
	if err != nil {
		log.Fatal("Error loading env file: ", err)
	}
}

var db *gorm.DB
var commitController *controllers.CommitController
var commitService *services.CommitService
var githubService *services.GithubService

// @title Documenting API (SavannahTech Task)
// @version 1
// @Description version 1 of api
// @contact.name Oladeji Oluwafemi
// @contact.url https://github.com/djfemz
// @contact.email oladejifemi00@gmail.com

// @host localhost:8080
// @BasePath /api/v1
func main() {
	configureAppComponents()
	go startFetchCommitsJob()
	go startFetchRepositoryMetaDataJob()
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/v1/commits", commitController.GetTopCommitAuthors)
	router.GET("/api/v1/commits/:repo", commitController.GetCommitsForRepository)
	router.GET("/api/v1/commits/since", commitController.GetCommitsByDateSince)
	port := os.Getenv("SERVER_PORT")
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server on port: ", port)
	}
}

func startFetchCommitsJob() {
	job := cron.New()
	_, err := job.AddFunc("* * */1 * *", func() {
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

func startFetchRepositoryMetaDataJob() {
	job := cron.New()
	_, err := job.AddFunc("* * */1 * *", func() {
		githubService.FetchRepositoryMetaData()
	})
	if err != nil {
		log.Println("Error creating job: ", err)
	}
	log.Println("Starting new task...")
	job.Start()
}

func configureAppComponents() {
	db, err := repositories.ConnectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to Datasource")
	}
	commitRepository := repositories.NewCommitRepository(db)
	githubAuxRepo := repositories.NewGithubAuxiliaryRepository(db)
	commitService = services.NewCommitService(commitRepository)
	githubAuxService := services.NewGithubRepoService(githubAuxRepo)
	githubService = services.NewGithubService(commitService, githubAuxService)
	commitController = controllers.NewCommitController(commitService)
}
