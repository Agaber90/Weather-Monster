package http

import (
	"Weather-Monster/models"
	"Weather-Monster/webhook"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

//WebhookHandler represent the httphandler for webhooks
type WebhookHandler struct {
	WebhoolUsecase webhook.WebhooksUsecase
}

//NewWebhookHandler will initialize the webhoooks/ resources endpoint
func NewWebhookHandler(e *echo.Echo, u webhook.WebhooksUsecase) {
	handler := &WebhookHandler{
		WebhoolUsecase: u,
	}
	e.POST("/webhoooks", handler.CreateWebhooks)
	e.DELETE("/webhoooks/:id", handler.DeleteWebhooks)
}

func (w *WebhookHandler) CreateWebhooks(e echo.Context) error {
	var webhooks models.Webhooks
	err := e.Bind(&webhooks)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = w.WebhoolUsecase.CreateWebhook(ctx, &webhooks)
	if err != nil {
		return e.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, webhooks)
}

func (w *WebhookHandler) DeleteWebhooks(e echo.Context) error {
	idP, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = w.WebhoolUsecase.DeleteWebhook(ctx, id)
	if err != nil {
		return e.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return e.NoContent(http.StatusNoContent)
}
func isRequestValid(m *models.Webhooks) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
