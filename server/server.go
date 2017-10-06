package server

import (
	//"fmt"
	//"github.com/VitaliiMichailovich/GGSMG/parser"
	//"github.com/VitaliiMichailovich/GGSMG/uri"
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/gin-gonic/contrib/static"
	"html/template"
)

var Router *gin.Engine

func IndexHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "IndexHandler",
		"text": ""})
}

func ProjectHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.HTML(http.StatusOK, "pages.html", gin.H{
		"title":    "About Project",
		"subtitle": "About Project",
		"text":     "<p>This project is a simple Golang+Gin project which is a result of my studying in IT Academy. </p> <p> What is it doing? Its parsing an index page in your web site, looking for a links in html. </p> <p> If it's the same like a domane it's and it'll check this link later. If it's contains '@', another domain or file extension it's ignoring. </p> <p> After checking all links It's has all links of your web-site. Next step it's creating a 'Sitemap.xml' file and send it to user's e-mail. </p>"})
}

func AboutMeHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.HTML(http.StatusOK, "pages.html", gin.H{
		"title":    "About Me",
		"subtitle": "About Me",
		"text":     template.HTML("<p> I just trying to improve my GoLang skills. </p>")})
}

func ContactHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	text := "<p> admin@micro.pp.ua </p>"
	c.HTML(http.StatusOK, "pages.html", gin.H{
		"title":    "Contact",
		"subtitle": "Contact",
		"text":     text})
}

func PostHandler(c *gin.Context) {
	x, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Data(200, "text/html; charset=utf-8", []byte(err.Error()))
	}
	c.Data(200, "text/html; charset=utf-8", x[:])
}

func Server() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	Router = gin.Default()

	// Set up a static server.
	Router.Use(static.Serve("/client/", static.LocalFile("./client", true)))

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	Router.LoadHTMLGlob("templates/*")

	// Links handlers
	Router.GET("/", IndexHandler)
	Router.GET("/p/", ProjectHandler)
	Router.GET("/i/", AboutMeHandler)
	Router.GET("/c/", ContactHandler)
	Router.POST("/", PostHandler)

	// Start serving the application
	Router.Run()
}
