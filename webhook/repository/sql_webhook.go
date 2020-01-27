package repository

import (
	"Weather-Monster/models"
	"Weather-Monster/webhook"
	"context"
	"database/sql"
	"fmt"
)

type sqlWebhookRepository struct {
	Conn *sql.DB
}

//NewWebhookRepository will create an object that represent the webhook.WebhooksRepository interface
func NewWebhookRepository(Conn *sql.DB) webhook.WebhooksRepository {
	return &sqlWebhookRepository{Conn}
}

func (r *sqlWebhookRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Temperatures, error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {

		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	result := make([]*models.Temperatures, 0)
	for rows.Next() {
		t := new(models.Temperatures)

		err = rows.Scan(
			&t.CityID,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (r *sqlWebhookRepository) CreateWebhook(ctx context.Context, w *models.Webhooks) error {
	qry := "INSERT INTO webhoooks SET city_id=?,callback_url=?"
	stmt, err := r.Conn.PrepareContext(ctx, qry)
	if err != nil {
		return err
	}

	qres, err := stmt.ExecContext(ctx, w.CityID, w.CallbackURL)
	if err != nil {
		return err
	}
	lastID, err := qres.LastInsertId()
	if err != nil {
		return err
	}
	w.ID = lastID
	return nil
}

func (r *sqlWebhookRepository) DeleteWebhook(ctx context.Context, ID int64) error {
	qry := "DELETE FROM webhoooks WHERE id=?"
	stmt, err := r.Conn.PrepareContext(ctx, qry)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, ID)
	if err != nil {

		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		return err
	}

	return nil
}

func (r *sqlWebhookRepository) GetCityByID(ctx context.Context, ctyID int64) (t *models.Temperatures, e error) {
	qry := "SELECT city_id FROM temperatures WHERE city_id=?"
	lst, err := r.fetch(ctx, qry, ctyID)

	if err != nil {
		return nil, err
	}

	if len(lst) > 0 {
		t = lst[0]
	} else {
		return nil, models.ErrNotFound
	}
	return
}
