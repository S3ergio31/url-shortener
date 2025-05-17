package controllers

import (
	"github.com/S3ergio31/url-shortener/application"
	"github.com/S3ergio31/url-shortener/http"
	"github.com/S3ergio31/url-shortener/infrastructure"
)

func Status(request *http.Request) http.Response {
	code := request.Param("code")

	repository := infrastructure.BuildShortRepository()

	short, err := application.ShortFinder(code, repository)

	if err != nil {
		return http.ResponseNotFound()
	}

	return http.ResponseAccepted(ToShortStatsResponse(*short))
}
