package routes

import (
	"emailn/internal/contract"
	internalerros "emailn/internal/internal-erros"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	if err != nil {
		if errors.Is(err, internalerros.ErrInternal) {
			render.Status(r, 500)
			return
		} else {
			render.Status(r, 400)
		}

		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 202)
	render.JSON(w, r, map[string]string{"id": id})
}
