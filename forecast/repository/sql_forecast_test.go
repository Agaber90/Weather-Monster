package repository_test

import (
	"Weather-Monster/forecast/repository"
	"Weather-Monster/models"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestGetForecasts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockForecast := []models.Temperatures{
		models.Temperatures{
			CityID: 1,
			Min:    25,
			Max:    30,
			Sample: 5052,
		},
	}
	rows := sqlmock.NewRows([]string{"city_id", "min", "Max", "Sample"}).AddRow(mockForecast[0].CityID, mockForecast[0].Min, mockForecast[0].Max, mockForecast[0].Sample)
	qry := "SELECT city_id,max,min,sample FROM temperatures WHERE city_id=\\?"
	mock.ExpectQuery(qry).WillReturnRows(rows)
	f := repository.NewForecaseRepository(db)
	ctyID := int64(1)
	forecastLst, err := f.GetForecasts(context.TODO(), ctyID)
	assert.NoError(t, err)
	assert.NotNil(t, forecastLst)
}
