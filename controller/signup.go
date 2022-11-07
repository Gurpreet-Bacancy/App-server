package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (app *Application) Signup(c *fiber.Ctx) error {
	// POST request

	// Get user details from body

	// check if user already exist or not

	// hash password

	// store user details and hashed password in db

	// generate token

	// return token

	return c.SendString("signup successful!")
}
