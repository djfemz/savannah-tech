package main

import (
	"github.com/djfemz/savannahTechTask/api/controllers"
	"github.com/djfemz/savannahTechTask/api/repositories"
	routes "github.com/djfemz/savannahTechTask/api/router"
	"github.com/djfemz/savannahTechTask/api/services"
	"github.com/djfemz/savannahTechTask/api/utils"
	_ "github.com/djfemz/savannahTechTask/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var err error

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file: ", err)
	}
}

var db *gorm.DB
var commitController *controllers.CommitController
var repoController *controllers.RepoController
var commitService *services.CommitService
var repoDiscoveryService *services.RepoDiscoveryService
var commitManager *services.CommitManager
var commitMonitorService *services.CommitMonitorService

// @title Documenting API (SavannahTech Task)
// @version 1
// @Description version 1 of api
// @contact.name Oladeji Oluwafemi
// @contact.url https://github.com/djfemz
// @contact.email oladejifemi00@gmail.com

// @host localhost:8082
// @BasePath /api/v1
func main() {
	configureAppComponents()
	isGithubCredentialValid := strings.TrimSpace(os.Getenv("REPO_NAME")) != utils.EMPTY_STRING &&
		strings.TrimSpace(os.Getenv("REPO_OWNER")) != utils.EMPTY_STRING
	if isGithubCredentialValid {
		repoController.PullData()
	} else {
		log.Println("[WARN:]\t Repo name is empty, provide repository name to start pulling data")
	}
	router := gin.Default()
	routes.SetupRoutes(router, commitController, repoController)
	port := os.Getenv("SERVER_PORT")
	log.Println("port: ", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server on port: ", port)
	}
}

func configureAppComponents() {
	db, err = repositories.ConnectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to Datasource")
	}
	commitRepository := repositories.NewCommitRepository(db)
	githubAuxRepo := repositories.NewGithubAuxiliaryRepository(db)
	commitService = services.NewCommitService(commitRepository)
	githubAuxService := services.NewGithubRepoMetadataService(githubAuxRepo)
	repoDiscoveryService = services.NewRepoDiscoveryService(githubAuxService)
	commitManager = services.NewCommitManager(commitService, repoDiscoveryService)
	commitMonitorService = services.NewCommitMonitorService(commitManager)
	commitController = controllers.NewCommitController(commitService)
	repoController = controllers.NewRepoController(repoDiscoveryService, commitManager, commitMonitorService)
}
