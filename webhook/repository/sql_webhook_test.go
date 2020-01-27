package repository_test

import (
	"Weather-Monster/models"
	"Weather-Monster/webhook/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestGetCityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"city_id"}).AddRow(1)
	qry := "SELECT city_id FROM temperatures WHERE city_id=\\?"
	mock.ExpectQuery(qry).WillReturnRows(rows)

	w := repository.NewWebhookRepository(db)
	ctyID := int64(1)
	temp, err := w.GetCityByID(context.TODO(), ctyID)
	assert.NoError(t, err)
	assert.NotNil(t, temp)
}

func TestCreateWebhook(t *testing.T) {
	wbh := &models.Webhooks{
		CityID:      1,
		CallbackURL: "https://my.service.com/high-temperature",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := "INSERT INTO webhoooks SET city_id=\\?,callback_url=\\?"
	pre := mock.ExpectPrepare(qry)
	pre.ExpectExec().WithArgs(wbh.CityID, wbh.CallbackURL).WillReturnResult(sqlmock.NewResult(1, 1))

	w := repository.NewWebhookRepository(db)
	err = w.CreateWebhook(context.TODO(), wbh)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), wbh.ID)
}

func TestDeleteWebhook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	qry := "DELETE FROM webhoooks WHERE id=\\?"
	pre := mock.ExpectPrepare(qry)
	pre.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	w := repository.NewWebhookRepository(db)
	ctyID := int64(1)
	err = w.DeleteWebhook(context.TODO(), ctyID)
	assert.NoError(t, err)
}
