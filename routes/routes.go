package routes

import (
	"crypto/rsa"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializeRoutes(app *controller.Application, fibApp *fiber.App, privateKey *rsa.PrivateKey) {
	fibApp.Use(logger.New())
	fibApp.Get("/healthcheck", app.HealthCheck)
}
