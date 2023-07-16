package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrpthf1nd3r/kaizen-server/routes"
)

// StartServer configures and starts an HTTP server using
// Gin Gonic.
func StartServer() {
	ginEngine := gin.Default()

	rootRouter := ginEngine.Group("/")
	routes.EnableShowcaseRoute(rootRouter)

	log.Fatal(ginEngine.Run(":8080"))
}
