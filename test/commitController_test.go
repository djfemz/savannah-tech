package test

import (
	"github.com/djfemz/savannahTechTask/api/controllers"
	"github.com/djfemz/savannahTechTask/api/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTopNCommitAuthors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	commitRepository.On("FindTopCommitAuthors", mock.Anything).Return(loadTestAuthorData(), nil)
	var commitController = controllers.NewCommitController(services.NewCommitService(commitRepository))
	router := gin.Default()
	router.GET("/api/v1/commits", commitController.GetTopCommitAuthors)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/commits", nil)
	request.Header.Add("Content-Type", "application/json")
	query := request.URL.Query()
	query.Add("size", "3")
	headers := query.Encode()
	request.URL.RawQuery = headers
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotNil(t, writer.Body)
}

func TestShouldGetCommitsForRepository(t *testing.T) {
	gin.SetMode(gin.TestMode)
	commitRepository.On("FindCommitsForRepoByName", mock.Anything).Return(loadTestCommits(), nil)
	var commitController = controllers.NewCommitController(services.NewCommitService(commitRepository))
	router := gin.Default()
	router.GET("/api/v1/commits/repo", commitController.GetCommitsForRepository)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/commits/repo", nil)
	request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotNil(t, writer.Body)
}
