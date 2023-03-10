package controllers

import (
	"absensi/configs"
	"absensi/helpers"
	"absensi/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register godoc
// @Summary Create User .
// @Description Create User .
// @Tags AUTH
// @Produce json
// @Param Body body models.FormRegister true "the body to register user"
// @Success 201 {object} models.FormRegister
// @Router /auth/register [post]
func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
	}

	query := configs.DB.Create(&user)
	if query.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", query.Error)
	}

	return helpers.ResponseJson(c, http.StatusCreated, true, "Registered succesfully", user)
}

// Login godoc
// @Summary Login User.
// @Description Login User.
// @Tags AUTH
// @Param Body body models.FormLogin true "the body to login user"
// @Produce json
// @Success 200 {object} models.FormLogin
// @Router /auth/token/ [post]
func Login(c echo.Context) error {
	var user models.User
	var request models.User

	if err := c.Bind(&request); err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
	}

	record := configs.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
	}

	credentialError := helpers.CheckPassword(request.Password, user.Password)
	if !credentialError {
		return helpers.ResponseJson(c, http.StatusUnauthorized, false, "Invalid Credential", nil)
	}

	tokenString, err := helpers.GenerateToken(int(user.ID))
	if err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
	}

	response := map[string]string{
		"token": tokenString,
	}

	return helpers.ResponseJson(c, http.StatusOK, true, "-", response)
}
