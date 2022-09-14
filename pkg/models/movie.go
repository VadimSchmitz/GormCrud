package models

type Movie struct {
	IMDb_id      string  `gorm:"column:IMDb_id;primaryKey" json:"imdb_id"`
	Title        string  `gorm:"Title" json:"title"`
	Rating       float64 `gorm:"Rating" json:"rating"`
	Year         int     `gorm:"Year" json:"year"`
	Plot_summary string  `Plot_summary:"Title" json:"summary"`
}
