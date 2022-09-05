package usecase

import (
	"be11/apiclean/features/user"
	"be11/apiclean/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []user.Core{{ID: 1, Name: "alterra", Email: "mail@alterra.id", Password: "qwerty", Phone: "0812345", Address: "Jakarta"}}

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("SelectAllData").Return(returnData, nil).Once()

		usecase := New(repo)
		resultData, err := usecase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, resultData[0].ID, returnData[0].ID)
		repo.AssertExpectations(t)
	})
}

func TestPostData(t *testing.T) {
	dataLayerMock := new(mocks.UserData)
	t.Run("Success Post Data", func(t *testing.T) {
		dataLayerMock.On("InsertData", mock.Anything).Return(1, nil).Once()

		dataInput := user.Core{Name: "alta", Email: "alta@alterra.id", Password: "qwerty"}
		usecase := New(dataLayerMock)
		result, err := usecase.PostData(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		dataLayerMock.AssertExpectations(t)
	})

	t.Run("failed. Name is empty", func(t *testing.T) {
		// dataLayerMock.On("InsertData", mock.Anything).Return(1, nil)
		dataInput := user.Core{Email: "alta@alterra.id", Password: "qwerty"}
		usecase := New(dataLayerMock)
		result, err := usecase.PostData(dataInput)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
		dataLayerMock.AssertExpectations(t)
	})
}
