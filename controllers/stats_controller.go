package controllers

import (
	"net/http"
	"watchwise_be/config"
	"watchwise_be/models"

	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	var totalMovies int64
	var totalReviews int64
	var avgRating float64
	var totalViews int64

	config.DB.Model(&models.Movie{}).Count(&totalMovies)
	config.DB.Model(&models.Review{}).Count(&totalReviews)
	config.DB.Model(&models.Movie{}).Select("AVG(rating)").Scan(&avgRating)
	config.DB.Model(&models.Movie{}).Select("SUM(view_count)").Scan(&totalViews)

	c.JSON(http.StatusOK, gin.H{
		"total_movies":   totalMovies,
		"total_reviews":  totalReviews,
		"average_rating": avgRating,
		"total_views":    totalViews,
	})
}
