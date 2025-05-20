package domain

import (
	"time"

	"github.com/google/uuid"
)

type ShortCreatorDto struct {
	Url string
}

type ShortUpdaterDto struct {
	ShortCode string
	Url       string
}

type Short struct {
	Id          uuid.UUID
	Url         string
	ShortCode   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AccessCount int
}

func (s *Short) Count() {
	s.AccessCount++
}

type ShortCode struct{}
