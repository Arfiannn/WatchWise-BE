package models

type Review struct {
	ID       int     `gorm:"column:id_reviews;primaryKey;autoIncrement" json:"id_reviews"`
	IDMovies int     `gorm:"column:id_movies" json:"id_movies"`
	UserName string  `gorm:"column:user_name" json:"user_name"`
	Rating   float32 `gorm:"column:rating" json:"rating"`
	Comment  string  `gorm:"column:comment" json:"comment"`
	Date     string  `gorm:"column:date" json:"date"`

	Movie Movie `gorm:"foreignKey:IDMovies;references:ID"`
}
