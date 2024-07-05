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
	Delete(id string) error
}

type ServiceImpl struct {
	Repository Repository
	SendMail   func(campaign *Campaign) error
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)

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
		return nil, internalerros.ProcessErrorToReturn(err)
	}

	return &contract.CampaignResponse{
		ID:        campaign.ID,
		Name:      campaign.Name,
		Content:   campaign.Content,
		Status:    campaign.Status,
		CreatedBy: campaign.CreatedBy,
	}, nil
}

func (s *ServiceImpl) Cancel(id string) error {
	campaign, err := s.getCampaignAndValidateIfStatusIsPending(id)

	if err != nil {
		return err
	}

	campaign.CancelCampaign()
	err = s.Repository.Update(campaign)

	if err != nil {
		println(err.Error())
		return internalerros.ErrInternal
	}

	return nil
}

func (s *ServiceImpl) Delete(id string) error {
	campaign, err := s.getCampaignAndValidateIfStatusIsPending(id)

	if err != nil {
		return err
	}

	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerros.ErrInternal
	}

	return nil
}

func (s *ServiceImpl) Start(id string) error {
	campaign, err := s.getCampaignAndValidateIfStatusIsPending(id)

	if err != nil {
		return err
	}

	err = s.SendMail(campaign)
	if err != nil {
		return internalerros.ErrInternal
	}

	campaign.DoneCampaign()
	err = s.Repository.Update(campaign)
	if err != nil {
		return internalerros.ErrInternal
	}

	return nil
}

func (s *ServiceImpl) getCampaignAndValidateIfStatusIsPending(id string) (*Campaign, error) {
	campaign, err := s.Repository.FindByID(id)

	if err != nil {
		return nil, internalerros.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return nil, errors.New("Campaign status invalid")
	}

	return campaign, nil
}
