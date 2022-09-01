package entities

import "time"

type WalletCore struct {
	ID        uint
	Jenis     string
	Nomor     string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WalletRequest struct {
	Jenis  string `json:"jenis" form:"jenis"`
	Nomor  string `json:"nomor" form:"nomor"`
	UserID uint   `json:"user_id" form:"user_id"`
}

type WalletResponse struct {
	Jenis  string `json:"jenis" form:"jenis"`
	Nomor  string `json:"nomor" form:"nomor"`
	UserID uint   `json:"user_id" form:"user_id"`
}

func RequestToCoreWallet(dataRequest WalletRequest) WalletCore {
	dataCore := WalletCore{
		Jenis:  dataRequest.Jenis,
		Nomor:  dataRequest.Nomor,
		UserID: dataRequest.UserID,
	}
	return dataCore
}
