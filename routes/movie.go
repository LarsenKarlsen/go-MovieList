package routes

import (
	"movieList/controllers"

	"github.com/gin-gonic/gin"
)

func AddMovieRoutes(rg *gin.RouterGroup) {
	movie := rg.Group("/movie")

	movie.GET("/:id", controllers.GetMovieById)

	movie.GET("/", controllers.GetAllMovie)

	movie.POST("/", controllers.AddNewMovie)

	movie.PUT("/:id", controllers.UpdateMovie)

	movie.DELETE("/:id", controllers.DeleteMovie)

}
