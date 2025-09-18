package models

type Review struct {
	ID       int     `gorm:"column:id_reviews;primaryKey" json:"id"`
	IDMovies  int     `json:"id_movies"`
	UserName string  `json:"user_name"`
	Rating   float32 `json:"rating"`
	Comment  string  `json:"comment"`
	Date     string  `json:"date"`
}
