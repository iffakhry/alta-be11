package entities

import "time"

type WalletCore struct {
	ID        uint
	Jenis     string
	Nomor     string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      UserCore
}

type WalletRequest struct {
	Jenis  string `json:"jenis" form:"jenis"`
	Nomor  string `json:"nomor" form:"nomor"`
	UserID uint   `json:"user_id" form:"user_id"`
}

type WalletResponse struct {
	Jenis  string       `json:"jenis" form:"jenis"`
	Nomor  string       `json:"nomor" form:"nomor"`
	UserID uint         `json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
}

func RequestToCoreWallet(dataRequest WalletRequest) WalletCore {
	dataCore := WalletCore{
		Jenis:  dataRequest.Jenis,
		Nomor:  dataRequest.Nomor,
		UserID: dataRequest.UserID,
	}
	return dataCore
}

func CoreToResponseWallet(dataCore WalletCore) WalletResponse {
	dataResponse := WalletResponse{
		Jenis:  dataCore.Jenis,
		Nomor:  dataCore.Nomor,
		UserID: dataCore.UserID,
		User: UserResponse{
			Name: dataCore.User.Name,
		},
	}
	return dataResponse
}

func CoreToResponseWalletList(dataCore []WalletCore) []WalletResponse {
	var dataResponse []WalletResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, CoreToResponseWallet(v))
	}
	return dataResponse
}
