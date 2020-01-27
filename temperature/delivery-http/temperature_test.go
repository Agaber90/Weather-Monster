package http_test

import (
	"Weather-Monster/models"
	_tempHTTP "Weather-Monster/temperature/delivery-http"
	"Weather-Monster/temperature/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateTemp(t *testing.T) {
	mockTemp := models.Temperatures{
		Min:       20,
		Max:       35,
		CityID:    1,
		Sample:    55,
		Timestamp: time.Now(),
	}

	tmpMockTemp := mockTemp
	tmpMockTemp.ID = 0
	mockUsecase := new(mocks.TempUsecase)
	j, err := json.Marshal(tmpMockTemp)
	assert.NoError(t, err)
	mockUsecase.On("CreateTemperature", mock.Anything, mock.AnythingOfType("*models.Temperatures")).Return(nil)
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/temperatures", strings.NewReader(string(j)))

	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/temperatures")

	handler := _tempHTTP.TempHandler{
		TempUsecase: mockUsecase,
	}
	err = handler.CreatetTemperature(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUsecase.AssertExpectations(t)
}
