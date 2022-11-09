package controller

import (
	"errors"
	"strconv"
	"time"

	"App-server/helper"

	"github.com/Gurpreet-Bacancy/bcl/types"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"

	"github.com/golang-jwt/jwt"

	"github.com/shamaton/msgpack"
)

// TODO Implement Signup API

// Login godoc
// @Summary check user login
// @Description user success to login then generate active token
// @Tags root
// @Accept */*
// @Produce octet-stream
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func (app *Application) Login(c *fiber.Ctx) error {
	// POST request
	// TODO Password checking via hash
	var (
		userLogin     types.UserLoginRequest
		loginResponse types.UserLoginResponse
		err           error
	)

	// Unmarshal data from request body
	err = msgpack.Unmarshal(c.Body(), &userLogin)
	if err != nil {
		return helper.HandleError(c, 400, err, "Invalid data, please provide proper data")
	}

	// Check if user exist or not From DB
	user, err := app.models.User.GetUserByEmail(userLogin.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.HandleError(c, 400, err, "invalid credentials")
		}
		return helper.HandleError(c, 500, err, "Something went wrong while getting user details")
	}

	//	Create the Claims
	claims := jwt.MapClaims{
		"id":    strconv.FormatUint(uint64(user.ID), 10),
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
