package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.GetAll()

	return campaigns, 200, err
}

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	campaign, err := h.CampaignService.GetById(id)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return campaign, http.StatusOK, nil
}
