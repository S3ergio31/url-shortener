package domain

import (
	"fmt"
)

type ShortNotFound struct {
	Key string
}

func (s ShortNotFound) Error() string {
	return fmt.Sprintf("Short with key: '%s' not found", s.Key)
}
