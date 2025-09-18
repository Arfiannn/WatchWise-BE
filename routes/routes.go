package routes

import (
	"watchwise_be/controllers"

	"github.com/gin-gonic/gin"
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

	// Stats
	r.GET("/stats", controllers.GetStats)

	return r
}
