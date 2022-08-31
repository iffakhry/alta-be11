package controllers

import (
	"be11/apimvc/entities"
	"be11/apimvc/helper"
	"be11/apimvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	result, err := repositories.SelectAllUser()
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to get Data")
	}
	var userResponseData []entities.UserResponse
	for _, v := range result {
		userResponseData = append(userResponseData, entities.UserResponse{
			Name:    v.Name,
			Email:   v.Email,
			Address: v.Address,
		})
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", userResponseData))
}

func CreateUserController(c echo.Context) error {
	userData := entities.UserRequest{}
	err := c.Bind(&userData)
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusBadRequest, "failed to bind data")
	}

	userCoreData := entities.UserCore{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: userData.Password,
		Phone:    userData.Phone,
		Address:  userData.Address,
	}

	row, errInsert := repositories.InsertUser(userCoreData)
	if errInsert != nil || row == 0 {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to insert data")
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("insert data success"))

}
