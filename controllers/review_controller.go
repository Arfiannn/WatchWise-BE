package controllers

import (
	"net/http"
	"strconv"
	"watchwise_be/config"
	"watchwise_be/models"

	"github.com/gin-gonic/gin"
)

func GetReviewsByMovie(c *gin.Context) {
	movieID := c.Param("id")
	var reviews []models.Review
	config.DB.Where("id_movies = ?", movieID).Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}

func CreateReview(c *gin.Context) {
	movieID := c.Param("id")
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review.IDMovies, _ = strconv.Atoi(movieID)
	config.DB.Create(&review)
	c.JSON(http.StatusCreated, review)
}

func GetAllReviews(c *gin.Context) {
	var reviews []models.Review
	if err := config.DB.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}
