package utils

import (
	"github.com/djfemz/savannahTechTask/api/appErrors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	RFC_TIME_SUFFIX     = "Z"
	REPO                = "repo"
	SIZE                = "size"
	DATABASE_PORT       = "DATABASE_PORT"
	SINCE               = "since"
	EMPTY_STRING        = ""
	ACCEPT_HEADER_VALUE = "application/vnd.github+json"
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

func GetCommitCount() string {
	req, err := http.NewRequest(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), nil)
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	query := req.URL.Query()
	query.Add("page", "1")
	query.Add("per_page", "1")
	req.URL.RawQuery = query.Encode()
	client := http.Client{}
	res, err := client.Do(req)
	log.Println("response: ", res)
	header := res.Header.Get("Link")
	if header == EMPTY_STRING {
		return "0"
	}
	parts := strings.Split(header, ",")
	part := parts[1]
	parts = strings.Split(part, "&")
	part = parts[0]
	parts = strings.Split(part, "=")
	part = parts[1]
	log.Println("part: ", part)
	return part
}
