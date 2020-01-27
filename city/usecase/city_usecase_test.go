package usecase_test

import (
	"Weather-Monster/city/mocks"
	"Weather-Monster/city/usecase"
	"Weather-Monster/models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestCreateCity(t *testing.T) {
	mockCityRepo := new(mocks.Repository)
	mockCity := models.Cities{
		Name:      "Berlin",
		Latitude:  52.520008,
		Longitude: 13.404954,
	}

	t.Run("success", func(t *testing.T) {
		tempMocCity := mockCity
		tempMocCity.ID = 0
		mockCityRepo.On("CreateCity", mock.Anything, mock.AnythingOfType("*models.Cities")).Return(nil).Once()
		u := usecase.NewCityUseCase(mockCityRepo, time.Second*20)
		err := u.CreateCity(context.TODO(), &tempMocCity)
		assert.NoError(t, err)
		mockCityRepo.AssertExpectations(t)
	})
}

func TestUpdateCity(t *testing.T) {
	mockCityRepo := new(mocks.Repository)
	mockCity := models.Cities{
		Name:      "Berlin",
		Latitude:  52.520008,
		Longitude: 13.404954,
	}
	t.Run("success", func(t *testing.T) {
		mockCityRepo.On("UpdateCity", mock.Anything, &mockCity).Once().Return(nil)
		u := usecase.NewCityUseCase(mockCityRepo, time.Second*20)
		err := u.UpdateCity(context.TODO(), &mockCity)
		assert.NoError(t, err)
		mockCityRepo.AssertExpectations(t)
	})
}

func TestDeleteCity(t *testing.T) {
	mockCityRepo := new(mocks.Repository)
	mockCity := models.Cities{
		Name:      "Berlin",
		Latitude:  52.520008,
		Longitude: 13.404954,
	}

	t.Run("success", func(t *testing.T) {
		mockCityRepo.On("DeleteCity", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()
		u := usecase.NewCityUseCase(mockCityRepo, time.Second*20)
		err := u.DeleteCity(context.TODO(), mockCity.ID)
		assert.NoError(t, err)
		mockCityRepo.AssertExpectations(t)
	})

	t.Run("error-happens-in-db", func(t *testing.T) {
		mockCityRepo.On("DeleteCity", mock.Anything, mock.AnythingOfType("int64")).Once().Return(nil, errors.New("Unexpected Error")).Once()
		u := usecase.NewCityUseCase(mockCityRepo, time.Second*20)
		err := u.DeleteCity(context.TODO(), mockCity.ID)
		assert.Error(t, err)
		mockCityRepo.AssertExpectations(t)
	})
}
