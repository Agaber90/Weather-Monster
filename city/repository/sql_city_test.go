package repository_test

import (
	"Weather-Monster/city/repository"
	"Weather-Monster/models"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestCreateCity(t *testing.T) {
	cty := &models.Cities{
		Name:      "Berlin",
		Latitude:  52.520008,
		Longitude: 13.404954,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := "INSERT INTO cities SET name=\\?,latitude=\\?,longitude=\\?"

	pre := mock.ExpectPrepare(qry)
	pre.ExpectExec().WithArgs(cty.Name, cty.Latitude, cty.Longitude).WillReturnResult(sqlmock.NewResult(15, 1))
	c := repository.NewCityRepository(db)
	err = c.CreateCity(context.TODO(), cty)
	assert.NoError(t, err)
	assert.Equal(t, int64(15), cty.ID)
}

func TestUpdateCity(t *testing.T) {
	cty := &models.Cities{
		ID:        15,
		Name:      "Berlin",
		Latitude:  52.520008,
		Longitude: 13.404954,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := "UPDATE cities SET name=\\?,latitude=\\?,longitude=\\? WHERE id=\\?"
	pre := mock.ExpectPrepare(qry)
	pre.ExpectExec().WithArgs(cty.Name, cty.Latitude, cty.Longitude, cty.ID).WillReturnResult(sqlmock.NewResult(15, 1))
	c := repository.NewCityRepository(db)
	err = c.UpdateCity(context.TODO(), cty)
	assert.NoError(t, err)
}

func TestDeleteCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := "DELETE FROM cities WHERE id =\\?"
	pre := mock.ExpectPrepare(qry)
	pre.ExpectExec().WithArgs(15).WillReturnResult(sqlmock.NewResult(15, 1))
	c := repository.NewCityRepository(db)
	ctyID := int64(15)
	err = c.DeleteCity(context.TODO(), ctyID)
	assert.NoError(t, err)
}
