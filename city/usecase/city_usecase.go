package usecase

import (
	"Weather-Monster/city"
	"Weather-Monster/models"
	"context"
	"time"
)

type citiesUseCase struct {
	cityRepo       city.CityRepository
	contextTimeout time.Duration
}

// NewCityUseCase will create new an city Usecase object representation of city.CitiyUseCase interface
func NewCityUseCase(c city.CityRepository, timeout time.Duration) city.CitiyUseCase {
	return &citiesUseCase{
		cityRepo:       c,
		contextTimeout: timeout,
	}
}

func (u *citiesUseCase) CreateCity(c context.Context, cty *models.Cities) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	err := u.cityRepo.CreateCity(ctx, cty)
	if err != nil {
		return err
	}
	return nil
}

func (u *citiesUseCase) UpdateCity(c context.Context, cty *models.Cities) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.cityRepo.UpdateCity(ctx, cty)
}

func (u *citiesUseCase) DeleteCity(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.cityRepo.DeleteCity(ctx, id)
}
