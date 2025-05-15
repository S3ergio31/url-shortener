package controllers

import (
	"github.com/S3ergio31/url-shortener/http"
)

func Update(request *http.Request) http.Response {
	id := request.Param("code")

	return http.ResponseOk(struct{ Id string }{Id: id})
}
