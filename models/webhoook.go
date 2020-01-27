package models

//Webhooks data model
type Webhooks struct {
	ID          int64  `json:"id"`
	CityID      int64  `json:"city_id"`
	CallbackURL string `json:"callback_url"`
}
