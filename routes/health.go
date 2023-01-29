package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHealthRoutes(rg *gin.RouterGroup) {
	health := rg.Group("/health")

	health.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Status": http.StatusOK})
	})
}
