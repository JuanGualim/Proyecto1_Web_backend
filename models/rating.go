package models

type Rating struct {
	ID       int `json:"id"`
	SeriesID int `json:"series_id"`
	Rating   int `json:"rating"`
}
