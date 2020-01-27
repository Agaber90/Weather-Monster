package usecase_test

import (
	"Weather-Monster/models"
	"Weather-Monster/webhook/mocks"
	"Weather-Monster/webhook/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWebhook(t *testing.T) {
	mockWebhookRepo := new(mocks.Repository)
	mockWebhook := models.Webhooks{
		CityID:      1,
		CallbackURL: "https://my.service.com/high-temperature",
	}

	t.Run("success", func(t *testing.T) {
		tempWebhook := mockWebhook
		tempWebhook.ID = 0
		mockWebhookRepo.On("GetCityByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, models.ErrNotFound).Once()
		mockWebhookRepo.On("CreateWebhook", mock.Anything, mock.AnythingOfType("*models.Webhooks")).Return(nil).Once()
		u := usecase.NewWebhookUsecase(mockWebhookRepo, time.Second*20)

		err := u.CreateWebhook(context.TODO(), &tempWebhook)
		assert.NoError(t, err)
		assert.Equal(t, mockWebhook.ID, tempWebhook.ID)
		mockWebhookRepo.AssertExpectations(t)
	})
}

func TestGetCityByID(t *testing.T) {
	mockWebhookRepo := new(mocks.Repository)
	mockCity := models.Temperatures{
		CityID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockWebhookRepo.On("GetCityByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCity, nil).Once()
		u := usecase.NewWebhookUsecase(mockWebhookRepo, time.Second*20)
		cty, err := u.GetCityByID(context.TODO(), mockCity.ID)
		assert.NoError(t, err)
		assert.NotNil(t, cty)
		mockWebhookRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockWebhookRepo.On("GetCityByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()
		u := usecase.NewWebhookUsecase(mockWebhookRepo, time.Second*20)
		cty, err := u.GetCityByID(context.TODO(), mockCity.ID)
		assert.Error(t, err)
		assert.Nil(t, cty)
		mockWebhookRepo.AssertExpectations(t)
	})
}

func TestDeleteWebhook(t *testing.T) {
	mockWebhookRepo := new(mocks.Repository)
	mockWebhook := models.Webhooks{
		CityID:      2,
		CallbackURL: "https://my.service.com/high-temperature",
	}

	t.Run("success", func(t *testing.T) {
		mockWebhookRepo.On("DeleteWebhook", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()
		u := usecase.NewWebhookUsecase(mockWebhookRepo, time.Second*20)
		err := u.DeleteWebhook(context.TODO(), mockWebhook.ID)
		assert.NoError(t, err)
		mockWebhookRepo.AssertExpectations(t)
	})

	t.Run("error-happens-in-db", func(t *testing.T) {
		mockWebhookRepo.On("DeleteWebhook", mock.Anything, mock.AnythingOfType("int64")).Once().Return(nil, errors.New("Unexpected Error")).Once()
		u := usecase.NewWebhookUsecase(mockWebhookRepo, time.Second*20)
		err := u.DeleteWebhook(context.TODO(), mockWebhook.ID)
		assert.Error(t, err)
		mockWebhookRepo.AssertExpectations(t)
	})
}
