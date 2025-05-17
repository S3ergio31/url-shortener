package domain

import (
	"log"
)

func CountVisit(event Event) {
	shortFound := event.(ShortFound)
	short, _ := shortFound.Repository.FindByShortCode(shortFound.ShortCode)
	short.Count()
	shortFound.Repository.Save(*short)

	log.Printf("CountVisit: Event ShortFound with ShortCode=%s was proccessed", shortFound.ShortCode)
}
