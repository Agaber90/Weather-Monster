package http

import (
	"Weather-Monster/city"
	"Weather-Monster/models"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

//CityHandler represent the httphandler for cities
type CityHandler struct {
	CtyUseCase city.CitiyUseCase
}

//NewCityHandler will initialize the cities/ resources endpoint
func NewCityHandler(e *echo.Echo, u city.CitiyUseCase) {
	handler := &CityHandler{
		CtyUseCase: u,
	}

	e.POST("/cities", handler.CreateCity)
	e.PATCH("/cities/:id", handler.UpdateCity)
	e.DELETE("/cities/:id", handler.DeleteCity)
}

//CreateCity will create a new city by given request body
func (c *CityHandler) CreateCity(e echo.Context) error {
	var city models.Cities
	err := e.Bind(&city)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&city); !ok {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = c.CtyUseCase.CreateCity(ctx, &city)
	if err != nil {
		return e.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, city)

}

// DeleteCity will delete city by given param
func (c *CityHandler) DeleteCity(e echo.Context) error {
	idP, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = c.CtyUseCase.DeleteCity(ctx, id)

	if err != nil {
		return e.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return e.NoContent(http.StatusNoContent)
}

//UpdateCity will update city by given param
func (c *CityHandler) UpdateCity(e echo.Context) error {
	var city models.Cities
	err := e.Bind(&city)
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&city); !ok {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	idP, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}
	id := int64(idP)
	city.ID = id

	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = c.CtyUseCase.UpdateCity(ctx, &city)
	if err != nil {
		return e.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, city)
}

func isRequestValid(m *models.Cities) (bool, error) {
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
