package controller

import (
	"time"

	"App-server/helper"

	"github.com/Gurpreet-Bacancy/bcl/types"
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt"

	"github.com/shamaton/msgpack"
)

// TODO Implement Signup API

// Login godoc
// @Summary check user login
// @Description user success to login then generate active token
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [post]
func (app *Application) Login(c *fiber.Ctx) error {
	// POST request
	var (
		userLogin     types.UserLoginRequest
		loginResponse types.UserLoginResponse
		err           error
	)
	msgpckHeader := c.Get("content-type")
	if msgpckHeader != "application/octet-stream" {
		return helper.HandleError(c, 400, nil, "Invalid messagepack request, Please provide messagepack request")
	}

	// Unmarshal data from request body
	err = msgpack.Unmarshal(c.Body(), &userLogin)
	if err != nil {
		return helper.HandleError(c, 400, err, "Invalid data, please provide proper data")
	}

	// Check if user exist or not From DB
	user, err := app.models.User.GetUserByEmail(userLogin.Email)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while getting user details")
	}

	//	Create the Claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 10).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(app.privateKey)
	if err != nil {
		return helper.HandleError(c, 500, err, "Something went wrong while getting user details")
	}

	// Create Response Data
	loginResponse.Token = t

	// Unmarshal data from request body
	response, err := msgpack.Marshal(loginResponse)
	if err != nil {
		return helper.HandleError(c, 400, err, "Something went wrong")
	}

	return c.Send(response)
}
