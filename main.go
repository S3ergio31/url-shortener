package main

import (
	"log"
	"net/http"

	"github.com/S3ergio31/url-shortener/controllers"
	"github.com/S3ergio31/url-shortener/routing"
	"github.com/joho/godotenv"
)

/**
 * POST /shorten
 * {
 *  "url": "https://www.example.com/some/long/url"
 * }
 *
 * {
 *  "id": "1",
 *  "url": "https://www.example.com/some/long/url",
 *  "shortCode": "abc123",
 *  "createdAt": "2021-09-01T12:00:00Z",
 *  "updatedAt": "2021-09-01T12:00:00Z"
 * }
 * status 201, 400
 *
 * GET /shorten/abc123
 * status 200, 404
 *
 * PUT /shorten/abc123
 * {
 * "url": "https://www.example.com/some/updated/url"
 * }
 * status 200,404,400
 *
 * DELETE /shorten/abc123
 * status 204, 404
 *
 * GET /shorten/abc123/stats
 * {
 * "id": "1",
 * "url": "https://www.example.com/some/long/url",
 * "shortCode": "abc123",
 * "createdAt": "2021-09-01T12:00:00Z",
 * "updatedAt": "2021-09-01T12:00:00Z",
 * "accessCount": 10
 * }
 *status: 202, 404
 *
 *
 * status 200, 404
 */
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
