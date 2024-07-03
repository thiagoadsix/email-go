package campaign

import (
	"emailn/internal/contract"
	internalerros "emailn/internal/internal-erros"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetAll() (*[]contract.CampaignResponse, error)
	GetById(id string) (*contract.CampaignResponse, error)
	Cancel(id string) error
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerros.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImpl) GetAll() (*[]contract.CampaignResponse, error) {
	campaigns, err := s.Repository.FindAll()
	if err != nil {
		return nil, internalerros.ErrInternal
	}

	campaignResponses := make([]contract.CampaignResponse, len(*campaigns))

	for i, campaign := range *campaigns {
		campaignResponses[i] = contract.CampaignResponse{
			ID:      campaign.ID,
			Name:    campaign.Name,
			Content: campaign.Content,
			Status:  campaign.Status,
		}
	}

	return &campaignResponses, nil
}

func (s *ServiceImpl) GetById(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.FindByID(id)

	if err != nil {
		println(err.Error())
		return nil, internalerros.ErrInternal
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}

func (s *ServiceImpl) Cancel(id string) error {
	campaign, err := s.Repository.FindByID(id)

	if err != nil {
		println(err.Error())
		return internalerros.ErrInternal
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.CancelCampaign()
	err = s.Repository.Update(campaign)

	if err != nil {
		println(err.Error())
		return internalerros.ErrInternal
	}

	return nil
}
