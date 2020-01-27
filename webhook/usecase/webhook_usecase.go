package usecase

import (
	"Weather-Monster/models"
	"Weather-Monster/webhook"
	"context"
	"time"
)

type webhookUsecase struct {
	webhookRepo    webhook.WebhooksUsecase
	contextTimeout time.Duration
}

// NewWebhookUsecase will create new an city Usecase object representation of city.CitiyUseCase interface
func NewWebhookUsecase(w webhook.WebhooksUsecase, timeout time.Duration) webhook.WebhooksUsecase {
	return &webhookUsecase{
		webhookRepo:    w,
		contextTimeout: timeout,
	}
}

func (u *webhookUsecase) CreateWebhook(ctx context.Context, w *models.Webhooks) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	existCity, _ := u.GetCityByID(ctx, w.CityID)

	if existCity == nil {
		return models.ErrConflict
	}

	err := u.webhookRepo.CreateWebhook(ctx, w)
	if err != nil {
		return err
	}
	return nil

}

func (u *webhookUsecase) DeleteWebhook(ctx context.Context, ID int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.webhookRepo.DeleteWebhook(ctx, ID)
}

func (u *webhookUsecase) GetCityByID(ctx context.Context, ctyID int64) (*models.Temperatures, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	res, err := u.webhookRepo.GetCityByID(ctx, ctyID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
