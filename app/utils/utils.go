package utils

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

var (
	GITHUB_COMMIT_URL      = os.Getenv("GITHUB_API_COMMIT_URL")
	GITHUB_REPOSITORY_URL  = os.Getenv("GITHUB_API_REPOSITORY_URL")
	GITHUB_REPOSITORY_NAME = os.Getenv("REPO_NAME")
	RFC_TIME_SUFFIX        = "Z"
)

func ExtractParamFromRequest(paramName string, ctx *gin.Context) (uint64, error) {
	return strconv.ParseUint(ctx.Query(paramName), 10, 64)
}
