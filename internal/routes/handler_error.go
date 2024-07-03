package routes

import (
	internalerros "emailn/internal/internal-erros"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type RouteFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(routeFunc RouteFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := routeFunc(w, r)

		if err != nil {
			if errors.Is(err, internalerros.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusBadRequest)
			}

			render.JSON(w, r, map[string]string{"error": err.Error()})

			return
		}

		render.Status(r, status)

		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
