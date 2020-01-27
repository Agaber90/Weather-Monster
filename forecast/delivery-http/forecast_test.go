package http_test

import (
	_forcastHTTP "Weather-Monster/forecast/delivery-http"
	"Weather-Monster/forecast/mocks"
	"Weather-Monster/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/bxcodec/faker"
)

func TestGetForecast(t *testing.T) {
	var mockForcast models.Temperatures
	err := faker.FakeData(&mockForcast)
	assert.NoError(t, err)
	mockUCase := new(mocks.Usecase)
	mockListForecast := make([]*models.Temperatures, 0)
	mockListForecast = append(mockListForecast, &mockForcast)
	ctyID := int(mockForcast.ID)
	mockUCase.On("GetForecasts", mock.Anything, int64(ctyID)).Return(mockListForecast, nil)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/forecasts/"+strconv.Itoa(ctyID), strings.NewReader(""))
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("forecasts/:city_id")
	c.SetParamNames("city_id")
	c.SetParamValues(strconv.Itoa(ctyID))

	handler := _forcastHTTP.ForecastHandler{
		ForecastUsecase: mockUCase,
	}
	err = handler.GetForecast(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
