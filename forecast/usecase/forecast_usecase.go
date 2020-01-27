package usecase

import (
	"Weather-Monster/forecast"
	"Weather-Monster/models"
	"context"
	"time"
)

type forecastUsecase struct {
	foreRepo       forecast.Forecastsusecase
	contextTimeout time.Duration
}

//NewForecastUsecase will create new an city Usecase object representation of forecast.Forecastsusecase interface
func NewForecastUsecase(f forecast.ForecastsRepository, timeout time.Duration) forecast.Forecastsusecase {
	return &forecastUsecase{
		foreRepo:       f,
		contextTimeout: timeout,
	}
}

func (u *forecastUsecase) GetForecasts(ctx context.Context, ctyID int64) ([]*models.Temperatures, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	foreLst, err := u.foreRepo.GetForecasts(ctx, ctyID)
	if err != nil {
		return nil, err
	}

	return foreLst, err

}
