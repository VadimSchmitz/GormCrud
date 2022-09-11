package models

type Movie struct {
	IMDb_id      string  `gorm:"column:IMDb_id;primaryKey" json:"id"`
	Title        string  `json:"name"`
	Year         int     `json:"year"`
	Rating       float64 `json:"score"`
	Plot_summary string  `json:"plot"`
}