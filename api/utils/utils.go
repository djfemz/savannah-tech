package utils

import (
	"github.com/djfemz/savannahTechTask/api/appErrors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	RFC_TIME_SUFFIX = "Z"
	REPO            = "repo"
	SIZE            = "size"
	DATABASE_PORT   = "DATABASE_PORT"
	SINCE           = "since"
)

func ExtractParamFromRequest(paramName string, ctx *gin.Context) (uint64, error) {
	return strconv.ParseUint(ctx.Query(paramName), 10, 64)
}

func GetTimeFrom(date string) (*time.Time, error) {
	log.Println(date)
	isoFormattedTime, err := time.Parse(os.Getenv("ISO_TIME_FORMAT"), date)
	if err != nil {
		log.Println("Date in wrong format: ", date)
		return nil, appErrors.NewTimeFormatError()
	}
	return &isoFormattedTime, nil
}
