package repository

import (
	"Weather-Monster/models"
	"Weather-Monster/temperature"

	"context"
	"database/sql"
)

type sqlTempRepository struct {
	Conn *sql.DB
}

//NewTempRepository will create an object that represent the temp.TempRepository interface
func NewTempRepository(Conn *sql.DB) temperature.TempRepository {
	return &sqlTempRepository{Conn}
}

func (r *sqlTempRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Cities, error) {
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

	result := make([]*models.Cities, 0)
	for rows.Next() {
		c := new(models.Cities)

		err = rows.Scan(
			&c.ID,
			&c.Name,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (r *sqlTempRepository) GetCityByID(ctx context.Context, ctyID int64) (c *models.Cities, e error) {
	qry := "SELECT id, name FROM weather.cities WHERE id=?"
	ctyLst, err := r.fetch(ctx, qry, ctyID)

	if err != nil {
		return nil, err
	}

	if len(ctyLst) > 0 {
		c = ctyLst[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (r *sqlTempRepository) CreateTemperature(ctx context.Context, t *models.Temperatures) error {
	qry := "INSERT INTO temperatures SET min=?,max=?,sample=?,city_id=?,timestamp=?"

	stmt, err := r.Conn.PrepareContext(ctx, qry)
	if err != nil {
		return err
	}

	qres, err := stmt.ExecContext(ctx, t.Min, t.Max, t.Sample, t.CityID, t.Timestamp)
	if err != nil {
		return err
	}

	lastID, err := qres.LastInsertId()
	if err != nil {
		return err
	}

	t.ID = lastID
	return nil
}
