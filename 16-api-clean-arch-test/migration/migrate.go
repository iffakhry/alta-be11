package migration

import (
	userModel "be11/apiclean/features/user/data"
	walletModel "be11/apiclean/features/wallet/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&walletModel.Wallet{})
}
