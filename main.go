package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/Nishad4140/url-shortner-golang/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/", routes.ShortenURL)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")
	}
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
