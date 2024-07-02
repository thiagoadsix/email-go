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
func (cr *CampaignRepository) FindAll() (*[]campaign.Campaign, error) {
	campaigns := make([]campaign.Campaign, len(cr.campaigns))
	copy(campaigns, cr.campaigns)
	return &campaigns, nil
}

func (cr *CampaignRepository) FindByID(id string) (*campaign.Campaign, error) {
	for _, c := range cr.campaigns {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, nil
}
