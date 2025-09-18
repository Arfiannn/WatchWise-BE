package controllers

import (
	"net/http"
	"strconv"
	"watchwise_be/config"
	"watchwise_be/models"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	config.DB.Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := config.DB.First(&movie, "id_movies = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&movie)
	c.JSON(http.StatusCreated, movie)
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := config.DB.First(&movie, "id_movies = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID, _ = strconv.Atoi(id)
	config.DB.Save(&movie)
	c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Movie{}, "id_movies = ?", id)
	c.Status(http.StatusNoContent)
}
