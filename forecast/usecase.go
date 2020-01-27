package forecast

import (
	"Weather-Monster/models"
	"context"
)

//Forecastsusecase intilaize the methods that will be implented to the usecase
type Forecastsusecase interface {
	GetForecasts(ctx context.Context, ctyID int64) ([]*models.Temperatures, error)
}
