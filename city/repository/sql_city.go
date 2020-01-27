package repository

import (
	"Weather-Monster/city"
	"Weather-Monster/models"
	"context"
	"database/sql"
	"fmt"
)

type sqlCityRepository struct {
	Conn *sql.DB
}

//NewCityRepository will create an object that represent the city.CityRepository interface
func NewCityRepository(Conn *sql.DB) city.CityRepository {
	return &sqlCityRepository{Conn}
}

func (r *sqlCityRepository) CreateCity(ctx context.Context, c *models.Cities) error {
	qry := "INSERT INTO cities SET name=?,latitude=?,longitude=?"

	stmt, err := r.Conn.PrepareContext(ctx, qry)
	if err != nil {
		return err
	}

	qres, err := stmt.ExecContext(ctx, c.Name, c.Latitude, c.Longitude)
	if err != nil {
		return err
	}

	lastID, err := qres.LastInsertId()
	if err != nil {
		return err
	}

	c.ID = lastID
	return nil
}

func (r *sqlCityRepository) UpdateCity(ctx context.Context, c *models.Cities) error {
	qry := "UPDATE cities SET name=?,latitude=?,longitude=? WHERE id=?"
	stmt, err := r.Conn.PrepareContext(ctx, qry)
	if err != nil {
		return nil
	}
	qres, err := stmt.ExecContext(ctx, c.Name, c.Latitude, c.Longitude, c.ID)
	if err != nil {
		return err
	}

	affect, err := qres.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		err = fmt.Errorf("City wasn't updated successful %d", affect)

		return err
	}
	return nil
}
func (r *sqlCityRepository) DeleteCity(ctx context.Context, id int64) error {
	qry := "DELETE FROM cities WHERE id =?"
	stmt, err := r.Conn.PrepareContext(ctx, qry)
	if err != nil {
		return err
	}

	qres, err := stmt.ExecContext(ctx, id)
	if err != nil {

		return err
	}

	rowsAfected, err := qres.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("City wasn't updated successful: %d", rowsAfected)
		return err
	}

	return nil
}
