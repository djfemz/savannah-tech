package utils

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

var (
	GITHUB_REPOSITORY_NAME = os.Getenv("REPO_NAME")
	RFC_TIME_SUFFIX        = "Z"
	REPO                   = "repo"
	SIZE                   = "size"
	DATABASE_PORT          = "DATABASE_PORT"
	SINCE                  = "since"
)

func ExtractParamFromRequest(paramName string, ctx *gin.Context) (uint64, error) {
	return strconv.ParseUint(ctx.Query(paramName), 10, 64)
}
