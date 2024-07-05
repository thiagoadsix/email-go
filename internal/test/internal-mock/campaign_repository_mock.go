package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Create(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) FindAll() (*[]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(1).(*[]campaign.Campaign), args.Error(0)
}

func (r *CampaignRepositoryMock) FindByID(id string) (*campaign.Campaign, error) {
	args := r.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*campaign.Campaign), nil
}

func (r *CampaignRepositoryMock) Update(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) Delete(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) GetCampaignsToBeSent() ([]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(1).([]campaign.Campaign), args.Error(0)
}
