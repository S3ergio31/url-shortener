package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/S3ergio31/url-shortener/application"
	"github.com/S3ergio31/url-shortener/controllers"
	"github.com/S3ergio31/url-shortener/domain"
	"github.com/S3ergio31/url-shortener/infrastructure"
)

type TestRequest struct {
	body   []byte
	method string
	path   string
}

func (testRequest *TestRequest) send() *http.Response {
	req := httptest.NewRequest(testRequest.method, testRequest.path, bytes.NewReader(testRequest.body))
	rec := httptest.NewRecorder()

	RegisterRoutes()
	Handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	return res
}

func assertEquals(t *testing.T, expected any, current any) {
	if current != expected {
		t.Errorf("expected %d, current %d", expected, current)
	}
}

func TestCreateShort(t *testing.T) {
	url := "https://www.example.com/some/long/url"
	testRequest := TestRequest{
		body:   []byte(fmt.Sprintf(`{"url": "%s"}`, url)),
		method: http.MethodPost,
		path:   "/shorten",
	}

	res := testRequest.send()

	var body controllers.ShortResponse
	json.NewDecoder(res.Body).Decode(&body)

	assertEquals(t, url, body.Url)
	assertEquals(t, http.StatusCreated, res.StatusCode)
}

func TestCreateShortBadRequest(t *testing.T) {
	testRequest := TestRequest{
		method: http.MethodPost,
		path:   "/shorten",
	}

	res := testRequest.send()

	assertEquals(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetShort(t *testing.T) {
	short := application.ShortCreator(
		domain.ShortCreatorDto{Url: "https://www.example.com/some/long/url"},
		infrastructure.BuildInMemoryRepository(),
	)

	testRequest := TestRequest{
		method: http.MethodGet,
		path:   "/shorten/" + short.ShortCode,
	}

	res := testRequest.send()

	var body controllers.ShortResponse
	json.NewDecoder(res.Body).Decode(&body)

	assertEquals(t, short.ShortCode, body.ShortCode)
	assertEquals(t, http.StatusOK, res.StatusCode)
}

func TestGetShortNotFound(t *testing.T) {
	testRequest := TestRequest{
		method: http.MethodGet,
		path:   "/shorten/aaa000",
	}

	res := testRequest.send()

	assertEquals(t, http.StatusNotFound, res.StatusCode)
}

func TestDeleteShort(t *testing.T) {
	short := application.ShortCreator(
		domain.ShortCreatorDto{Url: "https://www.example.com/some/long/url"},
		infrastructure.BuildInMemoryRepository(),
	)

	testRequest := TestRequest{
		method: http.MethodDelete,
		path:   "/shorten/" + short.ShortCode,
	}

	res := testRequest.send()

	assertEquals(t, http.StatusNoContent, res.StatusCode)
}

func TestDeleteShortNotFound(t *testing.T) {
	testRequest := TestRequest{
		method: http.MethodDelete,
		path:   "/shorten/aaa000",
	}

	res := testRequest.send()

	assertEquals(t, http.StatusNotFound, res.StatusCode)
}

func TestUpdateShort(t *testing.T) {
	short := application.ShortCreator(
		domain.ShortCreatorDto{Url: "https://www.example.com/some/long/url"},
		infrastructure.BuildInMemoryRepository(),
	)

	newUrl := "https://www.example.com/some/long/url-updated"
	testRequest := TestRequest{
		method: http.MethodPut,
		path:   "/shorten/" + short.ShortCode,
		body:   []byte(fmt.Sprintf(`{"url": "%s"}`, newUrl)),
	}

	res := testRequest.send()
	var body controllers.ShortResponse
	json.NewDecoder(res.Body).Decode(&body)

	assertEquals(t, newUrl, body.Url)
	assertEquals(t, http.StatusOK, res.StatusCode)
}

func TestUpdateShortBadRequest(t *testing.T) {
	testRequest := TestRequest{
		method: http.MethodPut,
		path:   "/shorten/aaa000",
	}

	res := testRequest.send()

	assertEquals(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetShortStats(t *testing.T) {
	short := application.ShortCreator(
		domain.ShortCreatorDto{Url: "https://www.example.com/some/long/url"},
		infrastructure.BuildInMemoryRepository(),
	)

	testRequest := TestRequest{
		method: http.MethodGet,
		path:   fmt.Sprintf("/shorten/%s/stats", short.ShortCode),
	}

	res := testRequest.send()

	var body controllers.ShortStatsResponse
	json.NewDecoder(res.Body).Decode(&body)

	assertEquals(t, 0, body.AccessCount)
	assertEquals(t, http.StatusOK, res.StatusCode)
}

func TestGetShortStatsNotFound(t *testing.T) {
	testRequest := TestRequest{
		method: http.MethodGet,
		path:   "/shorten/aaa000/stats",
	}

	res := testRequest.send()

	assertEquals(t, http.StatusNotFound, res.StatusCode)
}
