package mock

import (
	"emailn/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (mock *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
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
	return nil, nil
}

func (mock *CampaignServiceMock) Cancel(id string) error {
	return nil
}

func (mock *CampaignServiceMock) Delete(id string) error {
	return nil
}
