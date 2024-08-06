package controllers

import (
	"github.com/djfemz/savannahTechTask/api/services"
	"github.com/djfemz/savannahTechTask/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommitController struct {
	*services.CommitService
}

func NewCommitController(commitService *services.CommitService) *CommitController {
	return &CommitController{commitService}
}

// GetTopCommitAuthors godoc
// @Summary      Get top N commit Authors
// @Description  Get top N commit Authors where N is a number
// @Tags         Authors
// @Accept       json
// @Produce      json
// @Param        size   query   int  true  "Number of Authors"
// @Success      200  {object}  dtos.AuthorResponse
// @Failure      400  {object}  app-errors.authorsNotFound
// @Failure      404  {object}  app-errors.authorsNotFound
// @Failure      500  {object}  app-errors.authorsNotFound
// @Router       /commits/authors/top [get]
func (commitController *CommitController) GetTopCommitAuthors(ctx *gin.Context) {
	size, _ := utils.ExtractParamFromRequest(utils.SIZE, ctx)
	commits, err := commitController.CommitService.GetTopCommitAuthors(int(size))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, commits)
}

// GetCommitsForRepository godoc
// @Summary      Get commits for repository
// @Description  Get all commits for repository
// @Tags         Commits
// @Accept       json
// @Produce      json
// @Param        repo   path    string  true  "Repository Name"
// @Success      200  {object}  dtos.CommitResponse
// @Failure      400  {object}  app-errors.commitNotFound
// @Failure      404  {object}  app-errors.commitNotFound
// @Failure      500  {object}  app-errors.commitNotFound
// @Router       /commits/{repo} [get]
func (commitController *CommitController) GetCommitsForRepository(ctx *gin.Context) {
	repo := ctx.Param(utils.REPO)
	commits, err := commitController.CommitService.GetCommitsForRepo(repo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, commits)
}

// GetCommitsByDateSince godoc
// @Summary      Get commits by date
// @Description  Get all commits starting from the date supplied in the format (MM-DD-YYYY)
// @Tags         Commits
// @Accept       json
// @Produce      json
// @Param        since   query    string  true  "date"
// @Success      200  {object}  dtos.CommitResponse
// @Failure      400  {object}  app-errors.commitNotFound
// @Failure      404  {object}  app-errors.commitNotFound
// @Failure      500  {object}  app-errors.commitNotFound
// @Router       /commits/since [get]
func (commitController *CommitController) GetCommitsByDateSince(ctx *gin.Context) {
	since := ctx.Query(utils.SINCE)
	commits, err := commitController.CommitService.GetCommitsByDateSince(since)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, commits)
}
