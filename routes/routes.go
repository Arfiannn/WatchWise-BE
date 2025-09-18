package routes

import (
	"github.com/gin-gonic/gin"
	"watchwise_be/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Movies
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovie)
	r.POST("/movies", controllers.CreateMovie)
	r.PUT("/movies/:id", controllers.UpdateMovie)
	r.DELETE("/movies/:id", controllers.DeleteMovie)

	// Reviews
	r.GET("/movies/:id/reviews", controllers.GetReviewsByMovie)
	r.POST("/movies/:id/reviews", controllers.CreateReview)

	return r
}
