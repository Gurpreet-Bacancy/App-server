package controller

import (
	"App-server/helper"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/Gurpreet-Bacancy/bcl/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shamaton/msgpack"
)

//TODO:
// Make logging more traceble.

// GetUserLocation godoc
// @Summary get user location.
// @Description it takes user token and fetch user location from db.
// @Tags root
// @Accept */*
// @Produce octet-stream
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (app *Application) GetUserLocation(c *fiber.Ctx) error {
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

	// marshal
	response, err := msgpack.Marshal(coordinates)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while unmarshal coordinates")
	}

	return c.Send(response)
}

// AddUserLocation godoc
// @Summary add user new location if exits location it updates.
// @Description it takes user token and fetch user location from db if exits otherwise creates new.
// @Tags root
// @Accept */*
// @Produce octet-stream
// @Success 200 {object} map[string]interface{}
// @Router / [post]
func (app *Application) AddUserLocation(c *fiber.Ctx) error {
	// POST request
	var (
		coordinate model.Coordinates
		err        error
	)

	msgpckHeader := c.Get("content-type")

	if msgpckHeader != "application/octet-stream" {
		return helper.HandleError(c, 400, err, "Invalid messagepack request, Please provide messagepack request")
	}

	userClaims := c.Locals("user").(*jwt.Token)
	claims := userClaims.Claims.(jwt.MapClaims)
	userID := claims["id"].(uint)

	// Check if user exist or not From DB
	_, err = app.models.User.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return helper.HandleError(c, 404, err, "User not found")
		}

		return helper.HandleError(c, 500, err, "Something went wrong while getting user details")
	}

	// Unmarshal data from request body
	err = msgpack.Unmarshal(c.Body(), &coordinate)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while unmarshal user coordinates")
	}

	coordinate.UserID = userID

	// Upsert
	// check if user location is available if available then update, if not then add.
	userCoordinates, err := app.models.Coordinates.GetUserLocation(uint(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			err = app.models.Coordinates.AddUserLocation(userID, coordinate)
			if err != nil {
				return helper.HandleError(c, 500, err, "Something went wrong adding user location")
			}
			return helper.HandleError(c, 500, err, "Something went wrong adding user location")
		}

		return helper.HandleError(c, 500, err, "Something went wrong adding user location")
	}

	// if exist then update user location
	err = app.models.Coordinates.UpdateUserLocation(userCoordinates.ID, coordinate)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong updating user location")
	}

	return c.SendString("User location added sucessfully.")
}

// UpdateUserLocation godoc
// @Summary updates user new location if exits location it updates.
// @Description it takes user token and fetch user location from db and updates to it.
// @Tags root
// @Accept */*
// @Produce octet-stream
// @Success 200 {object} map[string]interface{}
// @Router / [put]
func (app *Application) UpdateUserLocation(c *fiber.Ctx) error {
	// PUT request
	var (
		coordinate model.Coordinates
		err        error
	)

	msgpckHeader := c.Get("content-type")
	if msgpckHeader != "application/octet-stream" {
		return helper.HandleError(c, 400, err, "Invalid messagepack request, Please provide messagepack request")
	}

	userClaims := c.Locals("user").(*jwt.Token)
	claims := userClaims.Claims.(jwt.MapClaims)
	userID := claims["id"].(uint)

	// Unmarshal data from request body
	err = msgpack.Unmarshal(c.Body(), &coordinate)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while unmarshal user")
	}

	coordinate.UserID = userID

	userCoordinates, err := app.models.Coordinates.GetUserLocation(uint(userID))
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong in getting user location")
	}

	err = app.models.Coordinates.UpdateUserLocation(userCoordinates.ID, coordinate)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong updating user location")
	}

	return c.SendString("User location updated sucessfully.")
}
