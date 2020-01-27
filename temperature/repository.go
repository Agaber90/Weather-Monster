package temperature

import (
	"Weather-Monster/models"
	"context"
)

//TempRepository that will be implemented the crud operation on Repositpry folder
type TempRepository interface {
	CreateTemperature(ctx context.Context, t *models.Temperatures) error
	GetCityByID(ctx context.Context, ctyID int64) (*models.Cities, error)
}
