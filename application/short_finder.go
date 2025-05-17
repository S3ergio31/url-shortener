package application

import (
	"github.com/S3ergio31/url-shortener/domain"
)

func ShortFinder(shortCode string, repository domain.ShortRepository) (*domain.Short, *domain.ShortNotFound) {
	short, err := repository.FindByShortCode(shortCode)

	return short, err

}
