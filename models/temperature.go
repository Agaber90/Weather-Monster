package models

import "time"

//Temperatures data model
type Temperatures struct {
	ID        int64     `json:"id"`
	CityID    int64     `json:"cityId"`
	Min       int64     `json:"min"`
	Max       int64     `json:"max"`
	Sample    int64     `json:"sample"`
	Timestamp time.Time `json:"timestamp"`
}
