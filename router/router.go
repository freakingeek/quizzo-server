package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stattusx/quizzo-server/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api_v1.0", logger.New())

	api.Get("/quiz", handler.SendMockQuizzesToClient)
}
