package usecase

import (
	"Weather-Monster/models"
	"Weather-Monster/temperature"
	"context"
	"time"
)

type temperatureUsecase struct {
	TempRepo       temperature.TempRepository
	contextTimeout time.Duration
}

//NewTempUseCase temp will create new an city Usecase object representation of city.CitiyUseCase interface
func NewTempUseCase(temp temperature.TempRepository, timeout time.Duration) temperature.TempUseCase {
	return &temperatureUsecase{
		TempRepo:       temp,
		contextTimeout: timeout,
	}
}

func (u *temperatureUsecase) CreateTemperature(ctx context.Context, t *models.Temperatures) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	existCity, _ := u.GetCityByID(ctx, t.CityID)
	if existCity == nil {
		return models.ErrConflict
	}

	err := u.TempRepo.CreateTemperature(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (u *temperatureUsecase) GetCityByID(ctx context.Context, ctyID int64) (*models.Cities, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, err := u.TempRepo.GetCityByID(ctx, ctyID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
