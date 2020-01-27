package http_test

import (
	"Weather-Monster/models"
	_webhooksHTTP "Weather-Monster/webhook/delivery-http"
	"Weather-Monster/webhook/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateWebhooks(t *testing.T) {
	mockWebhooks := models.Webhooks{
		CityID:      1,
		CallbackURL: "https://my.service.com/high-temperature",
	}

	tmpMockWebhook := mockWebhooks
	tmpMockWebhook.ID = 0
	mockUCase := new(mocks.WebhookUsecase)

	j, err := json.Marshal(tmpMockWebhook)
	assert.NoError(t, err)

	mockUCase.On("CreateWebhook", mock.Anything, mock.AnythingOfType("*models.Webhooks")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/webhoooks", strings.NewReader(string(j)))

	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/webhoooks")
	handler := _webhooksHTTP.WebhookHandler{
		WebhoolUsecase: mockUCase,
	}
	err = handler.CreateWebhooks(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestDeleteWebhooks(t *testing.T) {
	var mockWebhook models.Webhooks
	err := faker.FakeData(&mockWebhook)
	assert.NoError(t, err)

	mockUCase := new(mocks.WebhookUsecase)
	ID := int(mockWebhook.ID)
	mockUCase.On("DeleteWebhook", mock.Anything, int64(ID)).Return(nil)
	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/webhoooks/"+strconv.Itoa(ID), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("webhoooks/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(ID))

	handler := _webhooksHTTP.WebhookHandler{
		WebhoolUsecase: mockUCase,
	}

	err = handler.DeleteWebhooks(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)

}
