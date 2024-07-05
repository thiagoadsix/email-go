package routes

import (
	"bytes"
	"context"
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup2(body *contract.NewCampaign, createdByExpected string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/campaigns", &buf)

	ctx := context.WithValue(req.Context(), "email", createdByExpected)
	req = req.WithContext(ctx)

	return req, res
}

func Test_CampaignsPost_ShouldCreateNewCampaign(t *testing.T) {
	assert := assert.New(t)
	createdByExpected := "test@email.com"
	body := contract.NewCampaign{
		Name:    "Test Campaign",
		Content: "Test Content",
		Emails:  []string{"test@email.com"},
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(newCampaign contract.NewCampaign) bool {
		if newCampaign.Name == body.Name && newCampaign.Content == body.Content && newCampaign.CreatedBy == createdByExpected {
			return true
		} else {
			return false
		}
	})).Return("123", nil)
	handler := Handler{CampaignService: service}

	req, res := setup2(&body, createdByExpected)

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

	req, res := setup2(&body, "test@test.com")

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
