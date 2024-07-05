package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/mail"
	"emailn/internal/infrastructure/repository"
	"log"
	"time"

	godotenv "github.com/joho/godotenv"
)

func main() {
	println("Started worker...")

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loading .env file")
	}

	db := repository.NewClient()
	campaignRepository := repository.CampaignRepository{Db: db}
	campaignService := campaign.ServiceImpl{
		Repository: &repository.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}

	for {
		campaigns, err := campaignRepository.GetCampaignsToBeSent()

		if err != nil {
			println(err.Error())
		}

		println("Campaigns to be sent: ", len(campaigns))

		for _, campaign := range campaigns {
			println("Sending campaign: ", campaign.ID)
			campaignService.SendMailAndUpdateStatus(&campaign)
		}

		time.Sleep(10 * time.Second)
	}

}
