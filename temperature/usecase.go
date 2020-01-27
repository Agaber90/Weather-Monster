package temperature

import (
	"Weather-Monster/models"
	"context"
)

//TempUseCase intilaize the methods that will be implented to the usecase
type TempUseCase interface {
	CreateTemperature(ctx context.Context, t *models.Temperatures) error
	GetCityByID(ctx context.Context, ctyID int64) (*models.Cities, error)
}
