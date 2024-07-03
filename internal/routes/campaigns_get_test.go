package routes

import (
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGetById_ShouldReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignResponse{
		ID:      "123",
		Name:    "Test Campaign",
		Content: "Test Content",
		Status:  "pending",
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/campaigns/123", nil)

	response, status, _ := handler.CampaignGetById(res, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaign.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaign.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignsGetById_ShouldInformErrorWhenExists(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/campaigns/123", nil)

	_, _, err := handler.CampaignGetById(res, req)

	assert.NotNil(err)
}
