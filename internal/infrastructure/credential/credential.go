package credential

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/golang-jwt/jwt/v5"
)

func ValidateToken(token string, ctx context.Context) (string, error) {
	token = strings.Replace(token, "Bearer ", "", 1)
	provider, err := oidc.NewProvider(ctx, os.Getenv("KEYCLOAK_URL"))

	if err != nil {
		return "", errors.New("Error to connect to provider")
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})

	_, err = verifier.Verify(ctx, token)

	if err != nil {
		return "", errors.New("Invalid token")
	}

	tokenJwt, _ := jwtgo.Parse(token, nil)
	claims := tokenJwt.Claims.(jwtgo.MapClaims)

	return claims["email"].(string), nil
}
