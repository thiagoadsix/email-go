package routes

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Auth_WhenAuthorizationIsMissing_ReturnError(t *testing.T) {
	assert := assert.New(t)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("should not call next handler")
	})

	handlerError := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Contains(res.Body.String(), "Unauthorized")
}

func Test_Auth_WhenAuthorizationIsInvalid_ReturnError(t *testing.T) {
	assert := assert.New(t)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("should not call next handler")
	})
	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return "", errors.New("Unauthorized")
	}

	handlerError := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer invalid-token")
	res := httptest.NewRecorder()

	handlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Contains(res.Body.String(), "Unauthorized")
}

func Test_Auth_WhenAuthorizationIsValid_CallNextHandler(t *testing.T) {
	assert := assert.New(t)
	emailExpected := "test@email.com"
	var email string

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email = r.Context().Value("email").(string)
	})
	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return emailExpected, nil
	}

	handlerError := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer valid-token")
	res := httptest.NewRecorder()

	handlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Equal(emailExpected, email)
}
