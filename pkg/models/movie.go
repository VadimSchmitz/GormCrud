package models

type Movie struct {
	Id      int     `json:"id" gorm:"primaryKey"`
	Imdb_id string  `json:"imdb_id"`
	Title   string  `json:"title"`
	Rating  float32 `json:"rating"`
	Year    int     `json:"year"`
}
