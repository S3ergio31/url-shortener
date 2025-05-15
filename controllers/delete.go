package controllers

import (
	"github.com/S3ergio31/url-shortener/http"
)

func Delete(request *http.Request) http.Response {
	return http.ResponseNoContent()
}
