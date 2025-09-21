package controllers

import (
	"encoding/base64"
	"io"
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
	title := c.PostForm("title")
	genre := c.PostForm("genre")
	year, _ := strconv.Atoi(c.PostForm("year"))
	rating, _ := strconv.ParseFloat(c.PostForm("rating"), 32)
	trailer := c.PostForm("trailer")
	synopsis := c.PostForm("synopsis")

	file, err := c.FormFile("poster")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Poster file is required"})
		return
	}

	opened, _ := file.Open()
	defer opened.Close()
	fileBytes, _ := io.ReadAll(opened)

	posterBase64 := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(fileBytes)

	movie := models.Movie{
		Title:     title,
		Genre:     genre,
		Year:      year,
		Rating:    float32(rating),
		Synopsis:  synopsis,
		Poster:    posterBase64,
		Trailer:   trailer,
		ViewCount: 0,
	}

	if err := config.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := config.DB.First(&movie, "id_movies = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if title := c.PostForm("title"); title != "" {
		movie.Title = title
	}
	if genre := c.PostForm("genre"); genre != "" {
		movie.Genre = genre
	}
	if yearStr := c.PostForm("year"); yearStr != "" {
		year, _ := strconv.Atoi(yearStr)
		movie.Year = year
	}
	if ratingStr := c.PostForm("rating"); ratingStr != "" {
		rating, _ := strconv.ParseFloat(ratingStr, 32)
		movie.Rating = float32(rating)
	}
	if synopsis := c.PostForm("synopsis"); synopsis != "" {
		movie.Synopsis = synopsis
	}

	if trailer := c.PostForm("trailer"); trailer != "" {
		movie.Trailer = trailer
	}

	if file, err := c.FormFile("poster"); err == nil {
		opened, _ := file.Open()
		defer opened.Close()
		fileBytes, _ := io.ReadAll(opened)
		posterBase64 := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(fileBytes)
		movie.Poster = posterBase64
	}

	if err := config.DB.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Movie{}, "id_movies = ?", id)
	c.Status(http.StatusNoContent)
}

func View(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	if err := config.DB.First(&movie, "id_movies = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	movie.ViewCount += 1

	if err := config.DB.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update view count"})
		return
	}

	c.JSON(http.StatusOK, movie)
}
