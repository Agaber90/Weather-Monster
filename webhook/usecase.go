package webhook

import (
	"Weather-Monster/models"
	"context"
)

//WebhooksUsecase intilaize the methods that will be implented to the usecase
type WebhooksUsecase interface {
	CreateWebhook(ctx context.Context, w *models.Webhooks) error
	DeleteWebhook(ctx context.Context, ID int64) error
	GetCityByID(ctx context.Context, ctyID int64) (*models.Temperatures, error)
}
