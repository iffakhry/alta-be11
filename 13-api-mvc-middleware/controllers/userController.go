package controllers

import (
	"be11/apimvc/entities"
	"be11/apimvc/helper"
	"be11/apimvc/middlewares"
	"be11/apimvc/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	idtoken := middlewares.ExtractToken(c)
	log.Println("idtoken", idtoken)
	if idtoken != 1 {
		return helper.FailedResponseHelper(c, http.StatusUnauthorized, "you don't have access to this feature")
	}
	result, err := repositories.SelectAllUser()
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to get Data")
	}
	// var userResponseData []entities.UserResponse
	// for _, v := range result {
	// 	userResponseData = append(userResponseData, entities.CoreToResponse(v))
	// }
	// return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", userResponseData))
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", entities.CoreToResponseList(result)))
}

func CreateUserController(c echo.Context) error {
	userData := entities.UserRequest{}
	err := c.Bind(&userData)
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusBadRequest, "failed to bind data")
	}

	// userCoreData := entities.UserCore{
	// 	Name:     userData.Name,
	// 	Email:    userData.Email,
	// 	Password: userData.Password,
	// 	Phone:    userData.Phone,
	// 	Address:  userData.Address,
	// }
	var userCoreData = entities.RequestToCore(userData)

	row, errInsert := repositories.InsertUser(userCoreData)
	if errInsert != nil || row == 0 {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to insert data")
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("insert data success"))

}

func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return helper.FailedResponseHelper(c, http.StatusBadRequest, "param must be number")
	}
	log.Println(id)
	result, err := repositories.SelectUserById(idConv)
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to get data")
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", entities.CoreToResponse(result)))
}

func ProfileUserController(c echo.Context) error {
	idtoken := middlewares.ExtractToken(c)
	result, err := repositories.SelectProfileUser(idtoken)
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to get data profile")
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data profile", entities.CoreToResponse(result)))
}

func UpdateUserController(c echo.Context) error {

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return helper.FailedResponseHelper(c, http.StatusBadRequest, "param must be number")
	}

	idtoken := middlewares.ExtractToken(c)
	if idtoken != idConv {
		return helper.FailedResponseHelper(c, http.StatusUnauthorized, "unauthorized")
	}
	var dataUpdate entities.UserRequest
	errBind := c.Bind(&dataUpdate)
	log.Println(dataUpdate)
	if errBind != nil {
		return helper.FailedResponseHelper(c, http.StatusBadRequest, "error bind data")
	}

	row, err := repositories.UpdateUser(idConv, entities.RequestToCore(dataUpdate))
	if err != nil || row == 0 {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to update data")
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update data"))
}
