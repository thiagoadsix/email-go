package mock

import (
	"emailn/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (m *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := m.Called(newCampaign)

	return args.String(0), args.Error(1)
}

func (r *CampaignServiceMock) GetById(id string) (*contract.CampaignResponse, error) {
	args := r.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}

func (m *CampaignServiceMock) GetAll() (*[]contract.CampaignResponse, error) {
	return nil, nil
}
