package controller

import (
	"App-server/helper"

	"github.com/Gurpreet-Bacancy/bcl/model"
	"github.com/gofiber/fiber/v2"
	"github.com/shamaton/msgpack"
)

// GetNearestUser godoc
// @Summary get Get Nearest User.
// @Description it give nearest 10 user.
// @Tags root
// @Accept */*
// @Produce octet-stream
// @Success 200 {object} map[string]interface{}
// @Router /v1/nearest/user [post]
func (app *Application) GetNearestUser(c *fiber.Ctx) error {
	var (
		err            error
		reqCoordinates model.Coordinates
	)

	// Unmarshal data from request body
	err = msgpack.Unmarshal(c.Body(), &reqCoordinates)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while unmarshal user coordinates")
	}

	// get 10 nearest users
	Usercoordinates, err := app.models.Coordinates.GetNearestUsers(reqCoordinates)
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
