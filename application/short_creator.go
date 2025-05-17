package application

import (
	"math/rand"
	"time"

	"github.com/S3ergio31/url-shortener/domain"
	"github.com/google/uuid"
)

func ShortCreator(shortDto domain.ShortCreatorDto, repository domain.ShortRepository) domain.Short {
	short := domain.Short{
		Id:        uuid.New(),
		Url:       shortDto.Url,
		ShortCode: createShortCode(repository),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repository.Save(short)

	return short
}

func createShortCode(repository domain.ShortRepository) string {
	for {
		code := buildRandomString("abcdefghijklmnopqrstuvwxyz", 3) + buildRandomString("0123456789", 3)
		short, _ := repository.FindByShortCode(code)
		if short == nil {
			return code
		}
	}
}

func buildRandomString(seed string, stringLen int) string {
	if len(seed) == 0 {
		return ""
	}

	result := make([]rune, stringLen)

	for i := range result {
		result[i] = rune(seed[rand.Intn(len(seed))])
	}

	return string(result)
}
