package main

import (
	"App-server/routes"

	"github.com/gofiber/fiber"
)

var app *controller.Application

func init() {

}

func main() {
	fibApp := fiber.New()

	// initialize routes
	routes.InitializeRoutes(app, fibApp, privateKey)

	// starting server
	fibApp.Listen(":3000")
}
