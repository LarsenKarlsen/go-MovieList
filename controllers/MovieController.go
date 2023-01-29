package controllers

import (
	"movieList/initializers"
	"movieList/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var body struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	ImageLink string `json:"image"`
}

func AddNewMovie(c *gin.Context) {
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
	}

	movie := models.Movie{Title: body.Title, Year: body.Year, ImageLink: body.ImageLink}
	result := initializers.DB.Create(&movie)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add movie to db",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"movie":  movie,
	})
}

func GetAllMovie(c *gin.Context) {
	var movies []models.Movie
	initializers.DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"movies": movies,
	})
}

func GetMovieById(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	initializers.DB.First(&movie, "id=?", id)

	if movie.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"movie":  "",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	initializers.DB.First(&movie, "id=?", id)

	if movie.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"movie":  "",
		})

		return
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	movie.ImageLink = body.ImageLink
	movie.Title = body.Title
	movie.Year = body.Year

	initializers.DB.Save(&movie)

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	initializers.DB.First(&movie, "id=?", id)

	if movie.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"movie":  "",
		})

		return
	}

	initializers.DB.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{
		"message": "movie deleted",
	})
}
