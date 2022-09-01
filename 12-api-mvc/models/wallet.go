package models

import (
	"be11/apimvc/entities"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Jenis  string `json:"jenis" form:"jenis"`
	Nomor  string `json:"nomor" form:"nomor"`
	UserID uint   `json:"user_id" form:"user_id"`
	User   User
}

func FromEntitiesWallet(dataCore entities.WalletCore) Wallet {
	dataModel := Wallet{
		Jenis:  dataCore.Jenis,
		Nomor:  dataCore.Nomor,
		UserID: dataCore.UserID,
	}
	return dataModel
}

func ToEntitiesWallet(dataModel Wallet) entities.WalletCore {
	dataEntities := entities.WalletCore{
		ID:        dataModel.ID,
		Jenis:     dataModel.Jenis,
		Nomor:     dataModel.Nomor,
		UserID:    dataModel.UserID,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		User: entities.UserCore{
			ID:      dataModel.User.ID,
			Name:    dataModel.User.Name,
			Email:   dataModel.User.Email,
			Address: dataModel.User.Address,
		},
	}

	return dataEntities
}

func ToEntitiesWalletList(dataModel []Wallet) []entities.WalletCore {
	var dataEntities []entities.WalletCore
	for _, v := range dataModel {
		dataEntities = append(dataEntities, ToEntitiesWallet(v))
	}
	return dataEntities
}
