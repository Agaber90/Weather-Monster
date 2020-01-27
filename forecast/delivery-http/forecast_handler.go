package http

import (
	"Weather-Monster/forecast"
	"Weather-Monster/models"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

//ForecastHandler represent the httphandler for forecast
type ForecastHandler struct {
	ForecastUsecase forecast.Forecastsusecase
}

//NewForecastHandler will initialize the forecasts/ resources endpoint
func NewForecastHandler(e *echo.Echo, u forecast.Forecastsusecase) {
	handler := &ForecastHandler{
		ForecastUsecase: u,
	}

	e.GET("/forecasts/:city_id", handler.GetForecast)

}

//GetForecast fetch forecasts by given city_id
func (f *ForecastHandler) GetForecast(e echo.Context) error {
	idp, err := strconv.Atoi(e.Param("city_id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	ctyID := int64(idp)
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	forecastLst, err := f.ForecastUsecase.GetForecasts(ctx, ctyID)
	if err != nil {
		return e.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return e.JSON(http.StatusOK, forecastLst)

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
