package repositories

import (
	"be11/apimvc/config"
	"be11/apimvc/entities"
	"be11/apimvc/models"
)

// melakukan insert data wallet ke table wallet
func InsertWallet(data entities.WalletCore) (row int, err error) {
	dataModel := models.FromEntitiesWallet(data)
	tx := config.DB.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func SelectAllWallet() ([]entities.WalletCore, error) {
	var data []models.Wallet
	// tx := config.DB.Find(&data)
	tx := config.DB.Preload("User").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dataCore := models.ToEntitiesWalletList(data)
	return dataCore, nil

}
