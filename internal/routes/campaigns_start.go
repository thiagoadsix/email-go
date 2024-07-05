package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignStart(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Start(id)

	return nil, http.StatusOK, err
}
