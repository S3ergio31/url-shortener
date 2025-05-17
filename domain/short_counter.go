package domain

import (
	"log"
)

func VisitCounter(event Event) {
	shortFound := event.(ShortFound)
	short, _ := shortFound.Repository.FindByShortCode(shortFound.ShortCode)
	short.Count()
	shortFound.Repository.Save(*short)

	log.Printf("VisitCounter: Event ShortFound with ShortCode=%s was proccessed", shortFound.ShortCode)
}
