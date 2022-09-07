package data

import (
	userModel "be11/apiclean/features/user/data"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Jenis  string
	Nomor  string
	UserID uint
	User   userModel.User
}
