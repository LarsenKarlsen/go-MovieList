package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var movies = []movie{
	{ID: "1", Title: "Movie 1: Chapter 1", Year: 2020},
	{ID: "2", Title: "Movie 2: Chapter 2", Year: 2021},
	{ID: "3", Title: "Movie 3: Chapter 3", Year: 2023},
}

// get all movies from db
func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

// add new Movie to list
func postMovie(c *gin.Context) {
	var newMovie movie

	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

// get Movie by ID
func getMovieByID(c *gin.Context) {
	id := c.Param("id")

	for _, m := range movies {
		if m.ID == id {
			c.IndentedJSON(http.StatusOK, m)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
}

func AddMovieRoutes(rg *gin.RouterGroup) {
	movies := rg.Group("/movie")

	movies.GET("/", getMovies)
	movies.GET("/:id", getMovieByID)
	movies.POST("/", postMovie)
}
