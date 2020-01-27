package forecast

import (
	"Weather-Monster/models"
	"context"
)

//ForecastsRepository that will be implemented the crud operation on Repositpry folder
type ForecastsRepository interface {
	GetForecasts(ctx context.Context, ctyID int64) ([]*models.Temperatures, error)
}
