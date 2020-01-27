package http_test

import (
	_cityHTTP "Weather-Monster/city/delivery-http"
	"Weather-Monster/city/mocks"
	"Weather-Monster/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"strings"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func TestCreateCity(t *testing.T) {
	mockCity := models.Cities{
		Name:      "Berlin",
		Latitude:  52.520008,
		Longitude: 13.404954,
	}

	tmpMockCity := mockCity
	tmpMockCity.ID = 0
	mockUCase := new(mocks.CityUsecase)

	j, err := json.Marshal(tmpMockCity)
	assert.NoError(t, err)

	mockUCase.On("CreateCity", mock.Anything, mock.AnythingOfType("*models.Cities")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/cities", strings.NewReader(string(j)))

	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cities")

	handler := _cityHTTP.CityHandler{
		CtyUseCase: mockUCase,
	}

	err = handler.CreateCity(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestDeleteCity(t *testing.T) {
	var mockCity models.Cities
	err := faker.FakeData(&mockCity)
	assert.NoError(t, err)

	mockUCase := new(mocks.CityUsecase)

	ctyID := int(mockCity.ID)
	mockUCase.On("DeleteCity", mock.Anything, int64(ctyID)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/cities/"+strconv.Itoa(ctyID), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("cities/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(ctyID))

	handler := _cityHTTP.CityHandler{
		CtyUseCase: mockUCase,
	}

	err = handler.DeleteCity(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateCity(t *testing.T) {
	mockCity := models.Cities{
		Name:      "Berlin",
		Latitude:  52.520008,
		Longitude: 13.404954,
	}

	ctyID := int(mockCity.ID)

	tmpMockCity := mockCity
	tmpMockCity.ID = 0
	mockUCase := new(mocks.CityUsecase)

	j, err := json.Marshal(tmpMockCity)
	assert.NoError(t, err)

	mockUCase.On("UpdateCity", mock.Anything, mock.AnythingOfType("*models.Cities")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PATCH, "/cities"+strconv.Itoa(ctyID), strings.NewReader(string(j)))

	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("cities/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(ctyID))

	handler := _cityHTTP.CityHandler{
		CtyUseCase: mockUCase,
	}

	err = handler.UpdateCity(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)

}
