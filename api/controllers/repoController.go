package controllers

import (
	dtos "github.com/djfemz/savannahTechTask/api/dtos/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type RepoController struct {
}

func NewRepoController() *RepoController {
	return &RepoController{}
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
		} else {
			ctx.JSON(http.StatusBadRequest, dtos.NewBaseResponse("repo sent successfully"))

		}
	}
}
