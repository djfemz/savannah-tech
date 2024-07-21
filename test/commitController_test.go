package test

import (
	"testing"
)

func TestGetTopNCommitAuthors(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	//t.Run("success", func(t *testing.T) {
	//	var commitController = controllers.NewCommitController(services.NewCommitService(mocks.NewCommitRepository(t)))
	//	router := gin.Default()
	//	router.POST("/api/v1/commits", commitController.GetTopCommitAuthors)
	//	writer := httptest.NewRecorder()
	//
	//	request, _ := http.NewRequest(http.MethodGet, "/api/v1/commits", nil)
	//	request.Header.Add("Content-Type", "application/json")
	//	router.ServeHTTP(writer, request)
	//	log.Println(writer.Body.String())
	//	assert.Equal(t, http.StatusOK, writer.Code)
	//	assert.NotNil(t, writer.Body)
	//})
}
