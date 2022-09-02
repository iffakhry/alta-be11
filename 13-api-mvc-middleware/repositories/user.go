package repositories

import (
	"be11/apimvc/config"
	"be11/apimvc/entities"
	"be11/apimvc/models"
	"errors"
)

func SelectAllUser() ([]entities.UserCore, error) {
	var data []models.User
	tx := config.DB.Find(&data)
	// tx := config.DB.Preload("Wallets").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore []entities.UserCore
	for _, v := range data {
		dataCore = append(dataCore, entities.UserCore{
			ID:       v.ID,
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
			Phone:    v.Phone,
			Address:  v.Address,
		})
	}
	return dataCore, nil
}

func InsertUser(data entities.UserCore) (int, error) {
	// userGormData := models.User{
	// 	Name:     data.Name,
	// 	Email:    data.Email,
	// 	Password: data.Password,
	// 	Phone:    data.Phone,
	// 	Address:  data.Address,
	// }

	userGormData := models.FromEntities(data)
	tx := config.DB.Create(&userGormData)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func SelectUserById(id int) (entities.UserCore, error) {
	var user models.User
	tx := config.DB.First(&user, id)
	if tx.Error != nil {
		return entities.UserCore{}, tx.Error
	}

	return models.ToEntities(user), nil
}

func SelectProfileUser(id int) (entities.UserCore, error) {
	var user models.User
	tx := config.DB.First(&user, id)
	if tx.Error != nil {
		return entities.UserCore{}, tx.Error
	}

	return models.ToEntities(user), nil
}

func UpdateUser(id int, dataUpdate entities.UserCore) (int, error) {
	// var user models.User
	// user.ID = uint(id)

	tx := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(models.FromEntities(dataUpdate))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return int(tx.RowsAffected), nil
}
