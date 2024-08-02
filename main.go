package main

import (
	"github.com/djfemz/savannahTechTask/api/controllers"
	"github.com/djfemz/savannahTechTask/api/repositories"
	"github.com/djfemz/savannahTechTask/api/services"
	"github.com/djfemz/savannahTechTask/api/utils"
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
var doneChannel chan bool
var errorChannel chan any

// @title Documenting API (SavannahTech Task)
// @version 1
// @Description version 1 of api
// @contact.name Oladeji Oluwafemi
// @contact.url https://github.com/djfemz
// @contact.email oladejifemi00@gmail.com

// @host localhost:8082
// @BasePath /api/v1
func main() {
	if os.Getenv("REPO_NAME") == utils.EMPTY_STRING {
		log.Println("Repo name is empty, provide repository name to start pulling data")
	} else {
		pullData()
	}
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/v1/commits/authors/top", commitController.GetTopCommitAuthors)
	router.GET("/api/v1/commits/:repo", commitController.GetCommitsForRepository)
	router.GET("/api/v1/commits/since", commitController.GetCommitsByDateSince)
	router.GET("/api/v1/repositories/:repo", repoController.AddRepoName)
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
	repoController = controllers.NewRepoController()
}

func pullData() {
	job := cron.New()
	configureAppComponents()
	doneChannel = make(chan bool)
	errorChannel = make(chan any)
	exists, _ := repoDiscoveryService.ExistsByName(os.Getenv("REPO_NAME"))
	if exists {
		log.Println("[INFO:]\tcommit monitor to start pulling data in 1 hour")
		id, err := job.AddFunc("@hourly", func() {
			log.Println("[INFO:]\tabout to start monitoring repo")
			commitMonitorService.StartJob()
		})
		if err != nil {
			log.Printf("[Error: starting task with id: %d. Failed with error: %v", id, err)
		}
		log.Println("[INFO:]\tregistered job to run hourly")
	} else {
		go repoDiscoveryService.StartJob(doneChannel, errorChannel)
		select {
		case status := <-doneChannel:
			log.Println("[INFO:]\tfinished fetching repository meta data", status)
			go commitManager.StartJob()
			break
		case errr := <-errorChannel:
			log.Println("[ERROR:] in pulldata:  failed to fetch repository data: ", errr)
			break
		}

	}

}
