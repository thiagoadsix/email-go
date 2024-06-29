package routes

import (
	internalerros "emailn/internal/internal-erros"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type RouteFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(routeFunc RouteFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := routeFunc(w, r)

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

		render.Status(r, status)

		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
