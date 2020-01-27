package webhook

import (
	"Weather-Monster/models"
	"context"
)

//WebhooksRepository that will be implemented the crud operation on Repositpry folder
type WebhooksRepository interface {
	CreateWebhook(ctx context.Context, w *models.Webhooks) error
	DeleteWebhook(ctx context.Context, ID int64) error
	GetCityByID(ctx context.Context, ctyID int64) (*models.Temperatures, error)
}
