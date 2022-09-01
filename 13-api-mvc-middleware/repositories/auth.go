package repositories

import (
	"be11/apimvc/config"
	"be11/apimvc/middlewares"
	"be11/apimvc/models"
	"errors"
)

func LoginUser(email, password string) (string, error) {
	var userData models.User
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&userData)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return "", errToken
	}
	return token, nil
}
