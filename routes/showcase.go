package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnableShowcaseRoute(router *gin.RouterGroup) {
	router.GET("/showcase", showcase)
}

func showcase(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}
