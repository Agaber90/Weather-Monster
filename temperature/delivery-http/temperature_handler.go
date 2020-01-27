package http

import (
	"Weather-Monster/models"
	"Weather-Monster/temperature"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

//TempHandler represent the httphandler for temperature
type TempHandler struct {
	TempUsecase temperature.TempUseCase
}

//NewTempHandler will initialize the temperature/ resources endpoint
func NewTempHandler(e *echo.Echo, u temperature.TempUseCase) {
	handler := &TempHandler{
		TempUsecase: u,
	}

	e.POST("/temperatures", handler.CreatetTemperature)
}

//CreatetTemperature will create a new temperature by given request body
func (t *TempHandler) CreatetTemperature(e echo.Context) error {
	var temp models.Temperatures
	temp.Sample = temp.Max + temp.Min
	temp.Timestamp = time.Now()
	err := e.Bind(&temp)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&temp); !ok {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = t.TempUsecase.CreateTemperature(ctx, &temp)
	if err != nil {
		return e.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return e.JSON(http.StatusCreated, temp)
}

func isRequestValid(m *models.Temperatures) (bool, error) {
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
	logrus.Error(err)
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
