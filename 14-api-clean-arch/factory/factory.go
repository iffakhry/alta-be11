package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userData "be11/apiclean/features/user/data"
	userDelivery "be11/apiclean/features/user/delivery"
	userUsecase "be11/apiclean/features/user/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

}
