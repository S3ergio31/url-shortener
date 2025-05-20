package infrastructure

import (
	"log"
	"os"

	"github.com/S3ergio31/url-shortener/domain"
)

func BuildShortRepository() domain.ShortRepository {
	if os.Getenv("REPOSITORY_DRIVER") == "memory" {
		log.Println("BuildShortRepository -> BuildInMemoryRepository")
		return BuildInMemoryRepository()
	}

	if os.Getenv("REPOSITORY_DRIVER") == "mysql" {
		log.Println("BuildShortRepository -> BuildMySqlRepository")
		return BuildMySqlRepository()
	}

	log.Println("BuildShortRepository -> Dafault repository driver 'memory'")
	return BuildInMemoryRepository()
}
