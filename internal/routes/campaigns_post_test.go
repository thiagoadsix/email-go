package routes

import (
	"bytes"
	"emailn/internal/contract"
	internalmock "emailn/internal/test/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsPost_ShouldSaveNewCampaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "Test Campaign",
		Content: "Test Content",
		Emails:  []string{"test@email.com"},
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(newCampaign contract.NewCampaign) bool {
		if newCampaign.Name == body.Name && newCampaign.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("123", nil)
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/campaigns", &buf)

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignsPost_ShouldInformErrorWhenExists(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "Test Campaign",
		Content: "Test Content",
		Emails:  []string{"test@email.com"},
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/campaigns", &buf)

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
