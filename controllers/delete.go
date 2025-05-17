package controllers

import (
	"github.com/S3ergio31/url-shortener/application"
	"github.com/S3ergio31/url-shortener/http"
	"github.com/S3ergio31/url-shortener/infrastructure"
)

func Delete(request *http.Request) http.Response {
	code := request.Param("code")

	err := application.ShortDeleter(
		code,
		infrastructure.BuildShortRepository(),
	)

	if err != nil {
		return http.ResponseNotFound()
	}

	return http.ResponseNoContent()
}
