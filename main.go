package main

import (
	"movieList/controllers"
	"movieList/initializers"
	"movieList/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	// serve static files
	r.Static("/static", "./static/assets")
	r.LoadHTMLGlob("views/*")

	frontendRoutes := []string{
		"/", "/login", "/logout", "/register", "signup", "main"}

	for _, route := range frontendRoutes {
		r.GET(route, controllers.Home)
	}

	api := r.Group("/api")
	routes.AddUserRoutes(api)
	routes.AddHealthRoutes(api)
	routes.AddMovieRoutes(api)

	r.Run()
}
