package routes

import (
	"crypto/rsa"

	"App-server/controller"

	_ "App-server/docs/fibersimple"

	"github.com/Gurpreet-Bacancy/bcl/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func InitializeRoutes(app *controller.Application, fibApp *fiber.App, privateKey *rsa.PrivateKey) {
	// Middleware
	fibApp.Use(logger.New())

	// Routes
	fibApp.Get("/healthcheck", app.HealthCheck)
	fibApp.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	// Routes
	fibApp.Post("/signup", app.Signup)

	// Routes
	fibApp.Post("/login", app.Login)

	v1 := fibApp.Group("/v1")

	// Middleware
	v1.Use(jwtware.New(jwtware.Config{
		KeyFunc: middleware.CustomKeyFunc(privateKey),
	}))

	// Routes
	v1.Post("/location", app.AddUserLocation)
	// Routes
	v1.Put("/location", app.UpdateUserLocation)
	// Routes
	v1.Get("/location", app.GetUserLocation)
	// Routes
	v1.Get("/nearest/user", app.GetNearestUser)

}
