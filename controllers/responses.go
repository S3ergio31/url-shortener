package controllers

import (
	"time"

	"github.com/S3ergio31/url-shortener/domain"
	"github.com/google/uuid"
)

type ShortResponse struct {
	Id        uuid.UUID `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ShortStatsResponse struct {
	ShortResponse
	AccessCount int `json:"access_count"`
}

func ToShortResponse(short domain.Short) ShortResponse {
	return ShortResponse{
		Id:        short.Id,
		Url:       short.Url,
		ShortCode: short.ShortCode,
		CreatedAt: short.CreatedAt,
		UpdatedAt: short.UpdatedAt,
	}
}

func ToShortStatsResponse(short domain.Short) ShortStatsResponse {
	return ShortStatsResponse{
		ShortResponse: ToShortResponse(short),
		AccessCount:   short.AccessCount,
	}
}
