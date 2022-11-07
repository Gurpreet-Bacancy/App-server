package controller

import "github.com/gofiber/fiber/v2"

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (app *Application) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hey There! Everything's working fine, Chill.")
}
