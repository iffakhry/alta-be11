package usecase

import "be11/apiclean/features/user"

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

	return 1, nil
}
