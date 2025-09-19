package models

type Movie struct {
	ID        int     `gorm:"column:id_movies;primaryKey;autoIncrement" json:"id_movies"`
	Title     string  `gorm:"column:title" json:"title"`
	Genre     string  `gorm:"column:genre" json:"genre"`
	Year      int     `gorm:"column:year" json:"year"`
	Rating    float32 `gorm:"column:rating" json:"rating"`
	Synopsis  string  `gorm:"column:synopsis" json:"synopsis"`
	Poster    string  `gorm:"column:poster" json:"poster"`
	ViewCount int     `gorm:"column:view_count" json:"view_count"`
}
