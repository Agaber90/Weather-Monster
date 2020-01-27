package usecase_test

import (
	"Weather-Monster/forecast/mocks"
	"Weather-Monster/forecast/usecase"
	"Weather-Monster/models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetForecasts(t *testing.T) {
	mockForeRepo := new(mocks.Repository)
	mockForecast := &models.Temperatures{
		CityID: 1,
		Min:    25,
		Max:    30,
		Sample: 2505,
	}

	mockLstForecast := make([]*models.Temperatures, 0)
	mockLstForecast = append(mockLstForecast, mockForecast)

	t.Run("success", func(t *testing.T) {
		mockForeRepo.On("GetForecasts", mock.Anything, mock.AnythingOfType("int64")).Return(mockLstForecast, 1).Once()
		u := usecase.NewForecastUsecase(mockForeRepo, time.Second*2)
		lst, err := u.GetForecasts(context.TODO(), 1)
		assert.Len(t, lst, len(mockLstForecast))
		assert.NoError(t, err)
		mockForeRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockForeRepo.On("GetForecasts", mock.Anything, mock.AnythingOfType("int64")).Return(mockLstForecast, nil).Once()
		u := usecase.NewForecastUsecase(mockForeRepo, time.Second*2)
		lst, err := u.GetForecasts(context.TODO(), 1)
		assert.Error(t, err)
		assert.Len(t, lst, len(mockLstForecast))
		mockForeRepo.AssertExpectations(t)
	})
}
