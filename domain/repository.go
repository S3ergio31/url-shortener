package domain

type ShortRepository interface {
	Save(short Short)
	Delete(shortCode string)
	FindByShortCode(shortCode string) (*Short, *ShortNotFound)
}
