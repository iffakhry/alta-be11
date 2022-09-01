package controllers

import (
	"be11/apimvc/entities"
	"be11/apimvc/helper"
	"be11/apimvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateWalletController(c echo.Context) error {
	walletData := entities.WalletRequest{}
	err := c.Bind(&walletData)
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
	var walletCoreData = entities.RequestToCoreWallet(walletData)

	row, errInsert := repositories.InsertWallet(walletCoreData)
	if errInsert != nil || row == 0 {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to insert data")
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("insert data success"))
}

func GetWalletController(c echo.Context) error {
	result, err := repositories.SelectAllWallet()
	if err != nil {
		return helper.FailedResponseHelper(c, http.StatusInternalServerError, "failed to get Data")
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", entities.CoreToResponseWalletList(result)))
}
