package entities

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Wallets   []WalletCore
}

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

type UserResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

func RequestToCore(dataRequest UserRequest) UserCore {
	dataCore := UserCore{
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Phone:    dataRequest.Phone,
		Address:  dataRequest.Address,
	}
	return dataCore
}

func CoreToResponse(dataCore UserCore) UserResponse {
	dataResponse := UserResponse{
		Name:    dataCore.Name,
		Email:   dataCore.Email,
		Address: dataCore.Address,
	}
	return dataResponse
}

func CoreToResponseList(dataCore []UserCore) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, CoreToResponse(v))
	}
	return dataResponse
}
