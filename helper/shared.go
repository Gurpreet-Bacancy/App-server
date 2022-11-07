package helper

import "github.com/gofiber/fiber/v2"

func HandleError(c *fiber.Ctx, code int, err error, msg string) error {
	return c.JSON(fiber.Map{
		"code":    code,
		"error":   err,
		"message": msg,
	})
}
