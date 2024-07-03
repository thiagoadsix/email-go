package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignCancel(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Cancel(id)

	return nil, http.StatusOK, err
}
