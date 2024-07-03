package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/repository"
	"emailn/internal/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := repository.NewClient()

	campaignService := campaign.ServiceImpl{
		Repository: &repository.CampaignRepository{Db: db},
	}
	handler := routes.Handler{
		CampaignService: &campaignService,
	}

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(routes.Auth)

		r.Get("/", routes.HandlerError(handler.CampaignGetAll))
		r.Get("/{id}", routes.HandlerError(handler.CampaignGetById))
		r.Patch("/cancel/{id}", routes.HandlerError(handler.CampaignCancel))
		r.Delete("/delete/{id}", routes.HandlerError(handler.CampaignDelete))
	})

	r.Post("/campaigns", routes.HandlerError(handler.CampaignPost))

	http.ListenAndServe(":3000", r)
}
