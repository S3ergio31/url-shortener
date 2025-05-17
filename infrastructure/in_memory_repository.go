package infrastructure

import "github.com/S3ergio31/url-shortener/domain"

type InMemoryRepository struct {
	shorts map[string]domain.Short
}

func (repo InMemoryRepository) Save(short domain.Short) {
	repo.shorts[short.ShortCode] = short
}

func (repo InMemoryRepository) FindByShortCode(shortCode string) (*domain.Short, *domain.ShortNotFound) {
	for _, short := range repo.shorts {
		if short.ShortCode == shortCode {
			return &short, nil
		}
	}

	return nil, &domain.ShortNotFound{Key: shortCode}
}

func (repo InMemoryRepository) Delete(shortCode string) {
	delete(repo.shorts, shortCode)
}

var repository *InMemoryRepository

func BuildInMemoryRepository() InMemoryRepository {
	if repository == nil {
		repository = &InMemoryRepository{shorts: make(map[string]domain.Short)}
	}

	return *repository
}
