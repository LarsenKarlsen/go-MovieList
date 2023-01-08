package main

import (
	"movieList/controllers"
	"movieList/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.Run()
}
