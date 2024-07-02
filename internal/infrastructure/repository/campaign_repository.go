package repository

import (
	"emailn/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := cr.Db.Create(campaign)
	return tx.Error
}
func (cr *CampaignRepository) FindAll() (*[]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := cr.Db.Find(&campaigns)
	return &campaigns, tx.Error
}

func (cr *CampaignRepository) FindByID(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := cr.Db.First(&campaign, "id = ?", id)
	return &campaign, tx.Error
}
