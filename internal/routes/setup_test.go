package routes

import (
	"context"
	internalmock "emailn/internal/test/internal-mock"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

var (
	service *internalmock.CampaignServiceMock
	handler = Handler{}
)

func setUp() {
	service = new(internalmock.CampaignServiceMock)
	handler.CampaignService = service
}

func newHttpTest(method string, url string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(method, url, nil)

	res := httptest.NewRecorder()

	return req, res
}

func addParameterToRequest(req *http.Request, keyParameter string, valueParameter string) *http.Request {
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add(keyParameter, valueParameter)
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
}
