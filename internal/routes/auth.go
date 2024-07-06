package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"emailn/internal/infrastructure/credential"
)

type ValidateTokenFunc func(token string, ctx context.Context) (string, error)

var ValidateToken ValidateTokenFunc = credential.ValidateToken

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Unauthorized"})
			return
		}

		email, err := ValidateToken(tokenString, r.Context())

		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Unauthorized"})
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
