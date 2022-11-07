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
