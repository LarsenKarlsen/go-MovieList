package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Movie struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Year  string `json:"year"`
}

func AddMovieRoutes(rg *gin.RouterGroup) {
	movie := rg.Group("/movie")

	movie.GET("/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Status": http.StatusOK})
	})

	movie.POST("/", func(ctx *gin.Context) {

	})
}
