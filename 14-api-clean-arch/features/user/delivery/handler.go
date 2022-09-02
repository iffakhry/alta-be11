package delivery

import (
	"be11/apiclean/features/user"
	"be11/apiclean/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	e.GET("/users", handler.GetAll)
}

func (delivery *UserDelivery) GetAll(c echo.Context) error {
	// c.Param("id")
	// c.QueryParam("page")
	// c.Bind()

	results, err := delivery.userUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", results))
}
