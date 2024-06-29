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

func (r *RepositoryMock) FindAll() ([]Campaign, error) {
	args := r.Called(campaign)
	return args.Get(1).([]Campaign), args.Error(0)
}

func (r *RepositoryMock) FindByID(id string) (struct {
	ID      string
	Name    string
	Content string
	Status  string
}, error) {
	args := r.Called(id)
	return args.Get(0).(struct {
		ID      string
		Name    string
		Content string
		Status  string
	}), args.Error(1)
}

var (
	campaign = contract.NewCampaign{
		Name:    "New Campaign",
		Content: "Content",
		Emails:  []string{"test1@email.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repository := new(RepositoryMock)
	repository.On("Save", mock.Anything).Return(nil)

	service.Repository = repository

	id, err := service.Create(campaign)

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
		if c.Name != campaign.Name ||
			c.Content != campaign.Content ||
			len(c.Contacts) != len(campaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repository
	service.Create(campaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repository := new(RepositoryMock)
	repository.On("Save", mock.Anything).Return(errors.New("error to save on repository"))

	service.Repository = repository
	_, err := service.Create(campaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_GetCampaignById(t *testing.T) {
	assert := assert.New(t)

	expectedCampaign := struct {
		ID      string
		Name    string
		Content string
		Status  string
	}{
		ID:      "123abc456efg",
		Name:    "Sample Campaign",
		Content: "Sample Content",
		Status:  Pending,
	}

	repository := new(RepositoryMock)
	repository.On("FindByID", "123abc456efg").Return(expectedCampaign, nil)

	service.Repository = repository
	result, err := service.GetById("123abc456efg")

	assert.Nil(err)
	assert.Equal(expectedCampaign, result)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositoryFindById(t *testing.T) {
	assert := assert.New(t)

	repository := new(RepositoryMock)
	repository.On("FindByID", mock.Anything).Return(struct {
		ID      string
		Name    string
		Content string
		Status  string
	}{}, errors.New("error to find by id on repository"))

	service.Repository = repository
	_, err := service.GetById("123abc456efg")

	assert.True(errors.Is(internalerros.ErrInternal, err))
}
