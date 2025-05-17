package application

import (
	"github.com/S3ergio31/url-shortener/domain"
)

func ShortDeleter(shortCode string, repository domain.ShortRepository) *domain.ShortNotFound {
	_, err := repository.FindByShortCode(shortCode)

	if err != nil {
		return err
	}

	repository.Delete(shortCode)

	return err
}
