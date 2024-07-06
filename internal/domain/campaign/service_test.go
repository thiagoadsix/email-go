package campaign_test

import (
	// "emailn/internal/campaign/contract"
	contract "emailn/internal/contract/campaign"
	"emailn/internal/domain/campaign"
	internalerros "emailn/internal/internal-erros"
	internalmock "emailn/internal/test/internal-mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaignRequest{
		Name:      "New Campaign",
		Content:   "Content",
		Emails:    []string{"test1@email.com"},
		CreatedBy: "test_created_by@email.com",
	}
	campaignPending *campaign.Campaign
	campaignStarted *campaign.Campaign
	repository      *internalmock.CampaignRepositoryMock
	service         = campaign.ServiceImpl{}
)

func setUp() {
	repository = new(internalmock.CampaignRepositoryMock)
	service = campaign.ServiceImpl{
		Repository: repository,
	}
	campaignPending, _ = campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	campaignStarted = &campaign.Campaign{ID: "123", Name: "Test Campaign", Content: "Test Content", Status: campaign.Started}
}

func setUpFindByIdRepository(campaign *campaign.Campaign) {
	repository.On("FindByID", mock.Anything).Return(campaign, nil)
}

func setUpUpdateRepository(err error) {
	repository.On("Update", mock.Anything).Return(err)
}

func setUpSendEmailService(err error) {
	sendMail := func(campaign *campaign.Campaign) error {
		return err
	}

	service.SendMail = sendMail
}

func Test_Create_Campaign(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repository.On("Create", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	setUp()
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaignRequest{})

	assert.False(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_CreateCampaign(t *testing.T) {
	setUp()
	repository.On("Create", mock.MatchedBy(func(c *campaign.Campaign) bool {
		if c.Name != newCampaign.Name ||
			c.Content != newCampaign.Content ||
			len(c.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositoryCreate(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repository.On("Create", mock.Anything).Return(errors.New("error to create on repository"))

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_GetCampaignById(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repository.On("FindByID", mock.MatchedBy(func(id string) bool {
		return id == campaignPending.ID
	})).Return(campaignPending, nil)

	result, err := service.GetById(campaignPending.ID)

	assert.Nil(err)
	assert.Equal(campaignPending.ID, result.ID)
	assert.Equal(campaignPending.Name, result.Name)
	assert.Equal(campaignPending.Content, result.Content)
	assert.Equal(campaignPending.Status, result.Status)
	assert.Equal(campaignPending.CreatedBy, result.CreatedBy)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositoryFindById(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repository.On("FindByID", mock.Anything).Return(nil, errors.New("error to find by id on repository"))

	_, err := service.GetById("123abc456efg")

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_Delete_ReturnRecordNotFoundWhenCampaignDoesNotExist(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignIdInvalid := "invalid"

	repository.On("FindByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_ReturnStatusInvalid_WhenCampaignStatusNotEqualsPending(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpFindByIdRepository(campaignStarted)

	err := service.Delete(campaignStarted.ID)

	assert.Equal("Campaign status invalid", err.Error())
}

func Test_Delete_ReturnInternalError_WhenDeleteHasProblem(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpFindByIdRepository(campaignPending)
	repository.On("Delete", mock.Anything).Return(errors.New("error to delete on repository"))

	err := service.Delete(campaignPending.ID)

	assert.Equal(internalerros.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNil_WhenDeleteHasSuccess(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpFindByIdRepository(campaignPending)
	repository.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign == campaignPending
	})).Return(nil)

	err := service.Delete(campaignPending.ID)

	assert.Nil(err)
}

func Test_Start_CampaignWasUpdated_StatusIsStarted(t *testing.T) {
	setUp()
	setUpSendEmailService(nil)
	setUpFindByIdRepository(campaignPending)

	repository.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return campaignPending.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Started
	})).Return(nil)

	service.Start(campaignPending.ID)

	assert.Equal(t, campaign.Started, campaignPending.Status)
}

func Test_Start_ReturnRecordNotFoundWhenCampaignDoesNotExist(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignIdInvalid := "invalid"

	repository.On("FindByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Start(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Start_ReturnStatusInvalid_WhenCampaignStatusNotEqualsPending(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpFindByIdRepository(campaignStarted)

	err := service.Start(campaignStarted.ID)

	assert.Equal("Campaign status invalid", err.Error())
}

func Test_SendMailAndUpdateStatus_WhenFail_StatusIsFail(t *testing.T) {
	setUp()
	setUpSendEmailService(errors.New("error to send email"))

	repository.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return campaignPending.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Fail
	})).Return(nil)

	service.SendMailAndUpdateStatus(campaignPending)

	repository.AssertExpectations(t)
}

func Test_SendMailAndUpdateStatus_WhenSuccess_StatusIsDone(t *testing.T) {
	setUp()
	setUpSendEmailService(nil)

	repository.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return campaignPending.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Done
	})).Return(nil)

	service.SendMailAndUpdateStatus(campaignPending)

	repository.AssertExpectations(t)
}
