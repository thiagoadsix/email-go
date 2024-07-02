package campaign

import (
	"emailn/internal/contract"
	internalerros "emailn/internal/internal-erros"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *RepositoryMock) FindAll() (*[]Campaign, error) {
	args := r.Called(newCampaign)
	return args.Get(1).(*[]Campaign), args.Error(0)
}

func (r *RepositoryMock) FindByID(id string) (*Campaign, error) {
	args := r.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Campaign), nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "New Campaign",
		Content: "Content",
		Emails:  []string{"test1@email.com"},
	}
	service = ServiceImpl{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repository := new(RepositoryMock)
	repository.On("Save", mock.Anything).Return(nil)

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

func Test_Create_SaveCampaign(t *testing.T) {
	repository := new(RepositoryMock)
	repository.On("Save", mock.MatchedBy(func(c *Campaign) bool {
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

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repository := new(RepositoryMock)
	repository.On("Save", mock.Anything).Return(errors.New("error to save on repository"))

	service.Repository = repository
	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_GetCampaignById(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repository := new(RepositoryMock)
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

	repository := new(RepositoryMock)
	repository.On("FindByID", mock.Anything).Return(nil, errors.New("error to find by id on repository"))

	service.Repository = repository
	_, err := service.GetById("123abc456efg")

	assert.True(errors.Is(internalerros.ErrInternal, err))
}
