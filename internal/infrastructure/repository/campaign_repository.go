package repository

import (
	"emailn/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	cr.campaigns = append(cr.campaigns, *campaign)
	return nil
}
func (cr *CampaignRepository) FindAll() ([]campaign.Campaign, error) {
	return cr.campaigns, nil
}

func (cr *CampaignRepository) FindByID(id string) (struct {
	ID      string
	Name    string
	Content string
	Status  string
}, error) {
	for _, c := range cr.campaigns {
		if c.ID == id {
			return struct {
				ID      string
				Name    string
				Content string
				Status  string
			}{
				ID:      c.ID,
				Name:    c.Name,
				Content: c.Content,
				Status:  c.Status,
			}, nil
		}
	}
	return struct {
		ID      string
		Name    string
		Content string
		Status  string
	}{}, nil
}
