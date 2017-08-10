package main

import (
	"github.com/VitaliiMichailovich/GGSMG/server"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	server.Router = gin.Default()

	server.Router.Use(static.Serve("client", static.LocalFile("client", false)))
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	server.Router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	server.Server()

	// Start serving the application
	server.Router.Run()
}
