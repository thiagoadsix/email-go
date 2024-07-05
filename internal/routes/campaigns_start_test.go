package routes

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsStart_200(t *testing.T) {
	setUp()
	campaignId := "1"

	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignId
	})).Return(nil)

	req, res := newHttpTest(http.MethodPatch, "/campaigns/start", nil)
	req = addParameterToRequest(req, "id", campaignId)

	_, status, err := handler.CampaignStart(res, req)

	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)
}

func Test_CampaignsStart_Error(t *testing.T) {
	setUp()
	errorExpected := errors.New("error")

	service.On("Start", mock.Anything).Return(errorExpected)

	req, res := newHttpTest(http.MethodPatch, "/campaigns/start", nil)

	_, _, err := handler.CampaignStart(res, req)

	assert.Equal(t, errorExpected, err)
}
