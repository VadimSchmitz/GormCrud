package models

type Omdb_body struct {
	IMDb_id      string  `json:"imdbID"`
	Title        string  `gorm:"Title" json:"Title"`
	Rating       float64 `gorm:"Rating" json:"imdbRating,string"`
	Year         int     `gorm:"Year" json:"Year,string"`
	Plot_summary string  `Plot_summary:"Title" json:"Plot"`
}