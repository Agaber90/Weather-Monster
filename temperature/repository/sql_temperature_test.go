package repository_test

import (
	"Weather-Monster/models"
	"Weather-Monster/temperature/repository"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestGetCityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	qry := "SELECT ID FROM cities WHERE id=?"
	mock.ExpectQuery(qry).WillReturnRows(rows)
	cty := repository.NewTempRepository(db)
	ctyID := int64(1)
	gtCtyID, err := cty.GetCityByID(context.TODO(), ctyID)
	assert.NoError(t, err)
	assert.NotNil(t, gtCtyID)

}

func TestCreateTemperature(t *testing.T) {
	timeStamp := time.Now()
	temp := &models.Temperatures{
		Min:       32,
		Max:       35,
		CityID:    1,
		Sample:    35 + 32,
		Timestamp: timeStamp,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := "INSERT INTO temperatures SET min=\\?,max=\\?,sample=\\?,city_id=\\?,timestamp=\\?"

	pre := mock.ExpectPrepare(qry)
	pre.ExpectExec().WithArgs(temp.Min, temp.Max, temp.Sample, temp.CityID, temp.Timestamp).WillReturnResult(sqlmock.NewResult(1, 1))
	tmp := repository.NewTempRepository(db)

	err = tmp.CreateTemperature(context.TODO(), temp)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), temp.ID)
}
