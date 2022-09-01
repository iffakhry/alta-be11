package controllers

import (
	"be11/apimvc/entities"
	"be11/apimvc/helper"
	"be11/apimvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var authData entities.AuthRequest
	err := c.Bind(&authData)
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusBadRequest, "failed to bind data")
	}
	token, errToken := repositories.LoginUser(authData.Email, authData.Password)
	if errToken != nil {
		return helper.FailedResponseHelper(c, http.StatusBadRequest, "login failed")
	}
	dataToken := map[string]interface{}{
		"token": token,
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("login success", dataToken))
}
