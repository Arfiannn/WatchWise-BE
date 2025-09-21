package routes

import (
	"time"
	"watchwise_be/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Movies
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovie)
	r.POST("/movies", controllers.CreateMovie)
	r.PUT("/movies/:id", controllers.UpdateMovie)
	r.PUT("/movies/:id/view", controllers.View)
	r.DELETE("/movies/:id", controllers.DeleteMovie)

	// Reviews
	r.GET("/movies/:id/reviews", controllers.GetReviewsByMovie)
	r.POST("/movies/:id/reviews", controllers.CreateReview)
	r.GET("/reviews", controllers.GetAllReviews)

	// Stats
	r.GET("/stats", controllers.GetStats)

	return r
}
