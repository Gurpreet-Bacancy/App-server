package controller

import (
	"strconv"

	"App-server/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shamaton/msgpack"
)

// GetNearestUser godoc
// @Summary get Get Nearest User.
// @Description it give nearest 10 user.
// @Tags root
// @Accept */*
// @Produce octet-stream
// @Success 200 {object} map[string]interface{}
// @Router /v1/nearest/user [get]
func (app *Application) GetNearestUser(c *fiber.Ctx) error {
	var err error

	userClaims := c.Locals("user").(*jwt.Token)
	claims := userClaims.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	intval, err := strconv.Atoi(userID)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while getting user coordinates details")
	}

	// if exist then update user location
	coordinates, err := app.models.Coordinates.GetUserLocation(uint(intval))
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while getting user coordinates details")
	}

	// get 10 nearest users
	Usercoordinates, err := app.models.Coordinates.GetNearestUsers(coordinates)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while getting user coordinates details")
	}

	// marshal
	response, err := msgpack.Marshal(Usercoordinates)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while unmarshal coordinates")
	}

	return c.Send(response)
}
