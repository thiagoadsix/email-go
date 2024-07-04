package repository

import (
	"emailn/internal/domain/campaign"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() *gorm.DB {
	dsn := os.Getenv("DATA_BASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error connecting to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
