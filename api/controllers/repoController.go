package controllers

import (
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/djfemz/savannahTechTask/api/services"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
)

type RepoController struct {
	*services.RepoDiscoveryService
	*services.CommitManager
	*services.CommitMonitorService
}

var doneChannel chan bool
var errorChannel chan any
var commitManagerDoneChannel chan bool

func NewRepoController(repoDiscoveryService *services.RepoDiscoveryService,
	commitManager *services.CommitManager,
	commitMonitorService *services.CommitMonitorService) *RepoController {
	return &RepoController{
		repoDiscoveryService,
		commitManager,
		commitMonitorService,
	}
}

// AddRepoName AddRepository godoc
// @Summary      Used to Add Repository to the application
// @Description  Used to Add Repository to the application
// @Tags         Repository
// @Accept       json
// @Produce      json
// @Param        repo   path   int  true  "Number of Authors"
// @Success      200  {object}  dtos.BaseResponse
// @Failure      400  {object}  dtos.BaseResponse
// @Router       /api/v1/repositories/:repo [get]
func (repoController *RepoController) AddRepoName(ctx *gin.Context) {
	repo := ctx.Param("repo")
	if repo != "" {
		err := os.Setenv("REPO_NAME", repo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.NewBaseResponse("failed to add repo name"))
			return
		} else {
			ctx.JSON(http.StatusOK, dtos.NewBaseResponse("repo sent successfully"))

		}
	}
}

func (repoController *RepoController) PullData() {
	job := cron.New()
	doneChannel = make(chan bool)
	errorChannel = make(chan any)
	commitManagerDoneChannel = make(chan bool)
	isExistingRepository, _ := repoController.ExistsByName(os.Getenv("REPO_NAME"))
	if !isExistingRepository {
		fetchFromRepo(repoController)
	}
	select {
	case <-doneChannel:
		monitorRepoForChanges(job, repoController)
		job.Start()
	}
}

func monitorRepoForChanges(job *cron.Cron, controller *RepoController) {
	log.Println("[INFO:]\tcommit monitor to start pulling data in 1 hour")
	id, err := job.AddFunc("*/2 * * * *", func() {
		log.Println("[INFO:]\tabout to start monitoring repo")
		go controller.CommitMonitorService.StartJob()
	})
	if err != nil {
		log.Printf("[Error: starting task with id: %d. Failed with error: %v", id, err)
	}
	log.Println("[INFO:]\tregistered job to run hourly")

}

func fetchFromRepo(controller *RepoController) {
	go controller.RepoDiscoveryService.StartJob(doneChannel, errorChannel)
	select {
	case status := <-doneChannel:
		log.Println("[INFO:]\tfinished fetching repository meta data", status)
		go controller.CommitManager.StartJob(&commitManagerDoneChannel)
		break
	case errr := <-errorChannel:
		log.Println("[ERROR:]\tfailed to fetch repository metadata: ", errr)
		break
	}
}
