package controllers

import (
	"time"

	"github.com/S3ergio31/url-shortener/http"
)

type CreateBody struct {
	Url string
}

type Short struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(request *http.Request) http.Response {
	var createBody CreateBody
	request.Body(&createBody)

	short := Short{
		Id:        1,
		Url:       createBody.Url,
		ShortCode: "abc123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// repository.save(short)

	return http.ResponseOk(short)
}
