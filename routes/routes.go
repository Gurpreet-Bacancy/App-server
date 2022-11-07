package routes

import (
	"crypto/rsa"

	"App-server/controller"

	"github.com/Gurpreet-Bacancy/bcl/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func InitializeRoutes(app *controller.Application, fibApp *fiber.App, privateKey *rsa.PrivateKey) {
	fibApp.Use(logger.New())
	fibApp.Get("/healthcheck", app.HealthCheck)
	fibApp.Post("/signin", app.Login)
	{
		v1 := fibApp.Group("/v1")

		v1.Use(jwtware.New(jwtware.Config{
			KeyFunc: middleware.CustomKeyFunc(privateKey),
		}))
		// v1.Post("/location", app.AddUserLocation)
		// v1.Put("/location", app.UpdateUserLocation)
		// v1.Get("/location", app.GetUserLocation)
	}
}
