package controllers

import (
	"github.com/S3ergio31/url-shortener/application"
	"github.com/S3ergio31/url-shortener/domain"
	"github.com/S3ergio31/url-shortener/http"
	"github.com/S3ergio31/url-shortener/infrastructure"
)

func Get(request *http.Request) http.Response {
	code := request.Param("code")

	repository := infrastructure.BuildShortRepository()

	short, err := application.ShortFinder(code, repository)

	if err != nil {
		return http.ResponseNotFound()
	}

	domain.Publish(domain.ShortFound{ShortCode: code, Repository: repository})

	return http.ResponseOk(ToShortResponse(*short))
}
