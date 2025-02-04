package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/mail"
	"emailn/internal/infrastructure/repository"
	"emailn/internal/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	godotenv "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := repository.NewClient()

	campaignService := campaign.ServiceImpl{
		Repository: &repository.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}
	handler := routes.Handler{
		CampaignService: &campaignService,
	}

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(routes.Auth)

		r.Post("/", routes.HandlerError(handler.CampaignPost))
		r.Get("/", routes.HandlerError(handler.CampaignGetAll))
		r.Get("/{id}", routes.HandlerError(handler.CampaignGetById))
		r.Patch("/cancel/{id}", routes.HandlerError(handler.CampaignCancel))
		r.Delete("/delete/{id}", routes.HandlerError(handler.CampaignDelete))
		r.Patch("/start/{id}", routes.HandlerError(handler.CampaignStart))
	})

	http.ListenAndServe(":3000", r)
}
