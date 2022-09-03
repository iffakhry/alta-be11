package data

import (
	"be11/apiclean/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) SelectAllData() ([]user.Core, error) {
	var data []User
	tx := repo.db.Find(&data)
	// tx := config.DB.Preload("Wallets").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// var dataCore []user.Core
	// for _, v := range data {
	// 	dataCore = append(dataCore, user.Core{
	// 		ID:       v.ID,
	// 		Name:     v.Name,
	// 		Email:    v.Email,
	// 		Password: v.Password,
	// 		Phone:    v.Phone,
	// 		Address:  v.Address,
	// 	})
	// }
	dataCore := toCoreList(data)
	return dataCore, nil
}

func (repo *userData) InsertData(data user.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
