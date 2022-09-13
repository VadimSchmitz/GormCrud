package models

type Movie struct {
	IMDb_id string  `gorm:"column:IMDb_id;primaryKey" json:"imdb_id"`
	Title   string  `json:"title"`
	Rating  float64 `json:"rating"`
	Year    int     `json:"year"`
}