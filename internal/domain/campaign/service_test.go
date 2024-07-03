package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerros "emailn/internal/internal-erros"
	internalmock "emailn/internal/test/internal-mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "New Campaign",
		Content: "Content",
		Emails:  []string{"test1@email.com"},
	}
	service = campaign.ServiceImpl{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repository := new(internalmock.CampaignRepositoryMock)
	repository.On("Create", mock.Anything).Return(nil)

	service.Repository = repository

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_CreateCampaign(t *testing.T) {
	repository := new(internalmock.CampaignRepositoryMock)
	repository.On("Create", mock.MatchedBy(func(c *campaign.Campaign) bool {
		if c.Name != newCampaign.Name ||
			c.Content != newCampaign.Content ||
			len(c.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repository
	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositoryCreate(t *testing.T) {
	assert := assert.New(t)

	repository := new(internalmock.CampaignRepositoryMock)
	repository.On("Create", mock.Anything).Return(errors.New("error to create on repository"))

	service.Repository = repository
	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_GetCampaignById(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repository := new(internalmock.CampaignRepositoryMock)
	repository.On("FindByID", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repository

	result, err := service.GetById(campaign.ID)

	assert.Nil(err)
	assert.Equal(campaign.ID, result.ID)
	assert.Equal(campaign.Name, result.Name)
	assert.Equal(campaign.Content, result.Content)
	assert.Equal(campaign.Status, result.Status)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositoryFindById(t *testing.T) {
	assert := assert.New(t)

	repository := new(internalmock.CampaignRepositoryMock)
	repository.On("FindByID", mock.Anything).Return(nil, errors.New("error to find by id on repository"))

	service.Repository = repository
	_, err := service.GetById("123abc456efg")

	assert.True(errors.Is(internalerros.ErrInternal, err))
}
