package internalmock

import (
	contract "emailn/internal/contract/campaign"
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (mock *CampaignServiceMock) Create(newCampaign contract.NewCampaignRequest) (string, error) {
	args := mock.Called(newCampaign)

	return args.String(0), args.Error(1)
}

func (mock *CampaignServiceMock) GetById(id string) (*contract.CampaignResponse, error) {
	args := mock.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}

func (mock *CampaignServiceMock) GetAll() (*[]contract.CampaignResponse, error) {
	args := mock.Called()

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*[]contract.CampaignResponse), args.Error(1)
}

func (mock *CampaignServiceMock) Cancel(id string) error {
	args := mock.Called(id)

	return args.Error(0)
}

func (mock *CampaignServiceMock) Delete(id string) error {
	args := mock.Called(id)

	return args.Error(0)
}

func (mock *CampaignServiceMock) Start(id string) error {
	args := mock.Called(id)

	return args.Error(0)
}

func (mock *CampaignServiceMock) SendMailAndUpdateStatus(campaign *campaign.Campaign) {
	mock.Called(campaign)
}
