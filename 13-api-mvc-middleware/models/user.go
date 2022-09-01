package models

import (
	"be11/apimvc/entities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Wallets  []Wallet
}

func FromEntities(dataCore entities.UserCore) User {
	dataModel := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
	}
	return dataModel
}

func ToEntities(dataModel User) entities.UserCore {
	dataEntities := entities.UserCore{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Phone:     dataModel.Phone,
		Address:   dataModel.Address,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}

	return dataEntities
}
