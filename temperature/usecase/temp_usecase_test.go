package usecase_test

import (
	"Weather-Monster/models"
	"Weather-Monster/temperature/mocks"
	"Weather-Monster/temperature/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCityByID(t *testing.T) {
	mockCityRepo := new(mocks.Repository)
	mockCity := models.Cities{
		Name: "Berlin",
	}

	t.Run("success", func(t *testing.T) {
		mockCityRepo.On("GetCityByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCity, nil).Once()
		u := usecase.NewTempUseCase(mockCityRepo, time.Second*2)
		cty, err := u.GetCityByID(context.TODO(), mockCity.ID)
		assert.NoError(t, err)
		assert.NotNil(t, cty)
		mockCityRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockCityRepo.On("GetCityByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()
		u := usecase.NewTempUseCase(mockCityRepo, time.Second*2)
		cty, err := u.GetCityByID(context.TODO(), mockCity.ID)
		assert.Error(t, err)
		assert.Nil(t, cty)
		mockCityRepo.AssertExpectations(t)
	})
}

func TestCreateTemperature(t *testing.T) {
	mockTemRep := new(mocks.Repository)
	mockTemp := models.Temperatures{
		Min:       30,
		Max:       25,
		CityID:    1,
		Sample:    2502,
		Timestamp: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tempMockTemp := mockTemp
		tempMockTemp.ID = 0
		mockTemRep.On("GetCityByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, models.ErrNotFound).Once()
		mockTemRep.On("CreateTemperature", mock.Anything, mock.AnythingOfType("*models.Temperatures")).Return(nil).Once()
		u := usecase.NewTempUseCase(mockTemRep, time.Second*2)

		err := u.CreateTemperature(context.TODO(), &tempMockTemp)

		assert.NoError(t, err)
		assert.Equal(t, mockTemp.ID, tempMockTemp.ID)
		mockTemRep.AssertExpectations(t)
	})

}
