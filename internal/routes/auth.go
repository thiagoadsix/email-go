package routes

import (
	"context"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
	jwtgo "github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Unauthorized"})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/emailn_realm")

		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": "Error to connect to provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})

		_, err = verifier.Verify(r.Context(), tokenString)

		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Invalid token"})
			return
		}

		token, _ := jwtgo.Parse(tokenString, nil)
		claims := token.Claims.(jwtgo.MapClaims)
		email := claims["email"]

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
