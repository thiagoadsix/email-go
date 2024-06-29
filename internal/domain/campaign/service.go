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

func (s *Service) GetAll() (*[]Campaign, error) {
	campaigns, err := s.Repository.FindAll()

	if err != nil {
		return nil, internalerros.ErrInternal
	}

	return &campaigns, nil
}

func (s *Service) GetById(id string) (*struct {
	ID      string
	Name    string
	Content string
	Status  string
}, error) {
	campaign, err := s.Repository.FindByID(id)

	if err != nil {
		return &struct {
			ID      string
			Name    string
			Content string
			Status  string
		}{}, internalerros.ErrInternal
	}

	return &campaign, nil
}
