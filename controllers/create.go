package controllers

import (
	"github.com/S3ergio31/url-shortener/application"
	"github.com/S3ergio31/url-shortener/domain"
	"github.com/S3ergio31/url-shortener/http"
	"github.com/S3ergio31/url-shortener/infrastructure"
)

func Create(request *http.Request) http.Response {
	var shortCreatorDto domain.ShortCreatorDto
	request.Body(&shortCreatorDto)

	short := application.ShortCreator(
		shortCreatorDto,
		infrastructure.BuildShortRepository(),
	)

	return http.ResponseCreated(ToShortResponse(short))
}
