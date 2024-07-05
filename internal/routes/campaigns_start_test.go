package routes

import (
	"context"
	internalmock "emailn/internal/test/internal-mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsStart_200(t *testing.T) {
	assert := assert.New(t)
	campaignId := "1"

	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == "1"
	})).Return(nil)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest(http.MethodPatch, "/campaigns/start", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", campaignId)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(res, req)

	assert.Equal(http.StatusOK, status)
	assert.Nil(err)
}

func Test_CampaignsStart_Error(t *testing.T) {
	assert := assert.New(t)
	errorExpected := errors.New("error")

	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.Anything).Return(errorExpected)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest(http.MethodPatch, "/campaigns/start", nil)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignStart(res, req)

	assert.Equal(errorExpected, err)
}
