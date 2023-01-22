package main

import (
	"movieList/controllers"
	"movieList/initializers"
	"movieList/middleware"

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

	r.POST("api/signup", controllers.Signup)
	r.POST("api/login", controllers.Login)
	r.GET("api/logout", controllers.Logout)
	r.GET("api/validate", middleware.RequireAuth, controllers.Validate)

	r.GET("api/me", middleware.RequireAuth, controllers.Me)

	r.Run()
}
