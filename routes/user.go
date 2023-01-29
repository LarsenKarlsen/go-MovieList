package routes

import (
	"movieList/controllers"
	"movieList/middleware"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	// add new User
	user.POST("/signup", controllers.Signup)
	// log in current user
	user.POST("/login", controllers.Login)
	// logout current user
	user.GET("/logout", controllers.Logout)
	// get current user
	user.GET("/me", middleware.RequireAuth, controllers.Me)
}
