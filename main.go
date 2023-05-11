package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stattusx/quizzo-server/router"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
