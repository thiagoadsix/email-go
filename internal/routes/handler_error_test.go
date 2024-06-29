package routes

import (
	internalerros "emailn/internal/internal-erros"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_WhenRouteReturnsInternalError(t *testing.T) {
	assert := assert.New(t)

	route := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 500, internalerros.ErrInternal
	}

	handlerError := HandlerError(route)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerros.ErrInternal.Error())
}

func Test_HandlerError_WhenRouteReturnsDomainError(t *testing.T) {
	assert := assert.New(t)

	route := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 400, errors.New("domain error")
	}

	handlerError := HandlerError(route)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), errors.New("domain error").Error())
}

func Test_HandlerError_WhenRouteReturnsObjAndStatus(t *testing.T) {
	assert := assert.New(t)

	type bodyForTest struct {
		Id int
	}

	objExpected := bodyForTest{Id: 1}

	route := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}

	handlerError := HandlerError(route)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handlerError.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	objReturned := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objExpected, objReturned)
}
