package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Wallets  []WalletModel
}

type WalletModel struct {
	gorm.Model
	Jenis  string
	Nomor  string
	UserID uint
	User   User
}
