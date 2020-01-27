package repository

import (
	"Weather-Monster/forecast"
	"Weather-Monster/models"
	"context"
	"database/sql"
)

type sqlForecastRepository struct {
	Conn *sql.DB
}

//NewForecaseRepository will create an object that represent the forecasr.TempRepository interface
func NewForecaseRepository(Conn *sql.DB) forecast.ForecastsRepository {
	return &sqlForecastRepository{Conn}
}

func (r *sqlForecastRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Temperatures, error) {
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
		f := new(models.Temperatures)

		err = rows.Scan(
			&f.CityID,
			&f.Max,
			&f.Min,
			&f.Sample,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, f)
	}

	return result, nil
}

func (r *sqlForecastRepository) GetForecasts(ctx context.Context, ctyID int64) ([]*models.Temperatures, error) {
	qry := "SELECT city_id,max,min,sample FROM temperatures WHERE city_id=?"
	forecstLst, err := r.fetch(ctx, qry, ctyID)
	if err != nil {
		return nil, err
	}
	return forecstLst, err
}
