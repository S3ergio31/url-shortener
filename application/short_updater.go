package application

import (
	"github.com/S3ergio31/url-shortener/domain"
)

func ShortUpdater(shortDto domain.ShortUpdaterDto, repository domain.ShortRepository) (*domain.Short, *domain.ShortNotFound) {
	short, err := repository.FindByShortCode(shortDto.ShortCode)

	if err != nil {
		return short, err
	}

	short.Url = shortDto.Url

	repository.Save(*short)

	return short, err
}
