package delivery

import "be11/apiclean/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

func toCore(data UserRequest) user.Core {
	return user.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
	}
}
