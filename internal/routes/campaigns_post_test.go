package routes

import (
	contract "emailn/internal/contract/campaign"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createdByExpected = "test@email.com"
)

func Test_CampaignsPost_201(t *testing.T) {
	setUp()
	body := contract.NewCampaignRequest{
		Name:    "Test Campaign",
		Content: "Test Content",
		Emails:  []string{"test@email.com"},
	}

	service.On("Create", mock.MatchedBy(func(newCampaign contract.NewCampaignRequest) bool {
		if newCampaign.Name == body.Name && newCampaign.Content == body.Content && newCampaign.CreatedBy == createdByExpected {
			return true
		} else {
			return false
		}
	})).Return("123", nil)

	req, res := newHttpTest(http.MethodPost, "/campaigns", body)
	req = addContextToRequest(req, "email", createdByExpected)

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(t, http.StatusCreated, status)
	assert.Nil(t, err)
}

func Test_CampaignsPost_Error(t *testing.T) {
	setUp()
	body := contract.NewCampaignRequest{
		Name:    "Test Campaign",
		Content: "Test Content",
		Emails:  []string{"test@email.com"},
	}

	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	req, res := newHttpTest(http.MethodPost, "/campaigns", &body)
	req = addContextToRequest(req, "email", createdByExpected)

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(t, err)
}
