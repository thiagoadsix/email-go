package repository

import (
	"emailn/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	cr.campaigns = append(cr.campaigns, *campaign)
	return errors.New("an error")
}
