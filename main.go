package main

import (
	"log"
	"net/http"

	"github.com/S3ergio31/url-shortener/routing"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	routing.RegisterRoutes()

	http.HandleFunc("/", routing.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Cannot load .env file")
	}
}
