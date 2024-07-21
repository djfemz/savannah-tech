package controllers

import (
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/djfemz/savannahTechTask/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommitController struct {
	*services.CommitService
}

func NewCommitController(commitService *services.CommitService) *CommitController {
	return &CommitController{commitService}
}

func (commitController *CommitController) GetTopCommitAuthors(ctx *gin.Context) {
	size, _ := utils.ExtractParamFromRequest("size", ctx)
	commits, err := commitController.CommitService.GetTopCommitAuthors(int(size))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, commits)
}
