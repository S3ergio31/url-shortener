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
	accessCount int
}

func (s *Short) Count() {
	s.accessCount++
}

func (s *Short) AccessCount() int {
	return s.accessCount
}

type ShortCode struct{}
