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
	}

	return dataEntities
}
