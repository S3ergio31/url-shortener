package main

import (
	"log"
	"net/http"

	"github.com/S3ergio31/url-shortener/controllers"
	"github.com/S3ergio31/url-shortener/routing"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	routing.Post("/shorten", controllers.Create)
	routing.Get("/shorten/:code", controllers.Get)
	routing.Put("/shorten/:code", controllers.Update)
	routing.Delete("/shorten/:code", controllers.Delete)
	routing.Get("/shorten/:code/stats", controllers.Status)

	http.HandleFunc("/", routing.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Cannot load .env file")
	}
}
