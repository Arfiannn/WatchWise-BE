package models

type Movie struct {
	ID        int     `gorm:"column:id_movies;primaryKey" json:"id"`
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Rating    float32 `json:"rating"`
	Synopsis  string  `json:"synopsis"`
	Poster    string  `json:"poster"`
	ViewCount int     `json:"view_count"`
}