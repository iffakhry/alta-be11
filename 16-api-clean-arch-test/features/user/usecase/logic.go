package usecase

import (
	"be11/apiclean/features/user"
	"errors"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) GetAll() ([]user.Core, error) {
	results, err := usecase.userData.SelectAllData()
	return results, err
}

func (usecase *userUsecase) PostData(data user.Core) (int, error) {
	if data.Name == "" {
		return -1, errors.New("nama tidak boleh kosong")
	}
	row, err := usecase.userData.InsertData(data)
	return row, err
}
