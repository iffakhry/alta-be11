package delivery

import "be11/apiclean/features/user"

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

func fromCore(data user.Core) UserResponse {
	return UserResponse{
		ID:      data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}
}

func fromCoreList(data []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
