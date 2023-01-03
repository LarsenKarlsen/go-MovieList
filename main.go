package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"movieList/routers"
)

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
}

func main() {
	router := gin.Default()

	router.Static("/client", "./client")

	router.LoadHTMLGlob("client/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"msg": "index page loaded"})
	})
	router.GET("/sign", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign.html", gin.H{"msg": "sign page loaded"})
	})

	router.GET("/health", health)

	api := router.Group("api")
	routers.AddMovieRoutes(api)
	routers.AddUserRoutes(api)

	router.Run("localhost:8080")
}
