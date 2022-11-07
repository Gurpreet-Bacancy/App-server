package controller

import "github.com/gofiber/fiber/v2"

func (app *Application) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hey There! Everything's working fine, Chill.")
}
