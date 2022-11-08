package controller

import (
	"fmt"
	"strconv"

	"App-server/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shamaton/msgpack"
)

// GetUserLocation godoc
// @Summary get user location.
// @Description it takes user token and fetch user location from db.
// @Tags root
// @Accept */*
// @Produce msgp
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (app *Application) GetNearestUser(c *fiber.Ctx) error {
	var err error

	msgpckHeader := c.Get("content-type")
	if msgpckHeader != "application/octet-stream" {
		return helper.HandleError(c, 400, err, "Invalid messagepack request, Please provide messagepack request")
	}

	userClaims := c.Locals("user").(*jwt.Token)
	claims := userClaims.Claims.(jwt.MapClaims)
	fmt.Println("claims----->", claims)
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
	Usercoordinates, err := app.models.Coordinates.GetNearestUsers(coordinates, 10)
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
