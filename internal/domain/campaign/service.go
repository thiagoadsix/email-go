package campaign

import (
	"emailn/internal/contract"
	internalerros "emailn/internal/internal-erros"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
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

func (s *Service) GetAll() ([]Campaign, error) {
	campaigns, err := s.Repository.FindAll()

	if err != nil {
		return nil, internalerros.ErrInternal
	}

	return campaigns, nil
}
