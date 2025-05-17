package controllers

import (
	"github.com/S3ergio31/url-shortener/application"
	"github.com/S3ergio31/url-shortener/domain"
	"github.com/S3ergio31/url-shortener/http"
	"github.com/S3ergio31/url-shortener/infrastructure"
)

func Update(request *http.Request) http.Response {
	var shortUpdaterDto domain.ShortUpdaterDto
	request.Body(&shortUpdaterDto)
	shortUpdaterDto.ShortCode = request.Param("code")

	if shortUpdaterDto.Url == "" {
		return http.ResponseBadRequest()
	}

	short, err := application.ShortUpdater(shortUpdaterDto, infrastructure.BuildShortRepository())

	if err != nil {
		return http.ResponseNotFound()
	}

	return http.ResponseOk(short)

}
