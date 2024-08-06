package routes

import (
	"github.com/djfemz/savannahTechTask/api/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, commitController *controllers.CommitController, repoController *controllers.RepoController) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/v1/commits/authors/top", commitController.GetTopCommitAuthors)
	router.GET("/api/v1/commits/:repo", commitController.GetCommitsForRepository)
	router.GET("/api/v1/commits/since", commitController.GetCommitsByDateSince)
	router.GET("/api/v1/repositories/:repo", repoController.AddRepoName)
}
