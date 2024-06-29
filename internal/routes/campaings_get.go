package routes

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.GetAll()

	return campaigns, 200, err
}
