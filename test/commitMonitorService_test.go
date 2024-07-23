package test

import (
	"github.com/djfemz/savannahTechTask/app/mocks"
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"os"
	"testing"
)

func TestFetchCommitData(t *testing.T) {
	commitRepository := new(mocks.CommitRepository)
	commitRepository.On("SaveAll", mock.Anything).Return(nil)
	commitMonitorService := services.NewCommitMonitorService(services.NewCommitService(commitRepository))
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, os.Getenv("GITHUB_API_COMMIT_URL"), func(request *http.Request) (*http.Response, error) {
		res, err := httpmock.NewJsonResponse(http.StatusOK, githubCommitResponse)
		return res, err
	})
	data, err := commitMonitorService.FetchCommitData(0)
	assert.Nil(t, err)
	assert.NotEmpty(t, data)
}
