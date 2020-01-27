package city

import (
	"Weather-Monster/models"
	"context"
)

//CityRepository that will be implemented the crud operation on Repositpry folder
type CityRepository interface {
	CreateCity(ctx context.Context, c *models.Cities) error
	UpdateCity(ctx context.Context, c *models.Cities) error
	DeleteCity(ctx context.Context, id int64) error
}
