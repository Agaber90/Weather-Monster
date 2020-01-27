package city

import (
	"Weather-Monster/models"
	"context"
)

//CitiyUseCase intilaize the methods that will be implented to the usecase
type CitiyUseCase interface {
	CreateCity(c context.Context, cty *models.Cities) error
	UpdateCity(c context.Context, cty *models.Cities) error
	DeleteCity(c context.Context, id int64) error
}
