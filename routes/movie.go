package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMovieRoutes(rg *gin.RouterGroup) {
	movie := rg.Group("/movie")

	movie.GET("/:id", func(ctx *gin.Context) {
		// get movie
		ctx.JSON(http.StatusOK, gin.H{"Status": http.StatusOK})
	})

	movie.GET("/", func(ctx *gin.Context) {
		// get all movies in db
		// TODO: make pagination
	})

	movie.POST("/", func(ctx *gin.Context) {
		// add new movie
	})

	movie.POST("/:id/edit", func(ctx *gin.Context) {
		// update movie
	})

	movie.POST("/:id/del", func(ctx *gin.Context) {
		// delete movie
	})

}
