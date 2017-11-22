package server

import (
	"fmt"
	"github.com/VitaliiMichailovich/GGSMG/checkIn"
	//"github.com/VitaliiMichailovich/GGSMG/sender"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"github.com/VitaliiMichailovich/GGSMG/parser"
	"github.com/VitaliiMichailovich/GGSMG/xmlworker"
)

var Router *gin.Engine

func IndexHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "IndexHandler",
		"text":  template.HTML("")})
}

func ProjectHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.HTML(http.StatusOK, "pages.html", gin.H{
		"title":    "About Project",
		"subtitle": "About Project",
		"text":     template.HTML("<p>This project is a simple Golang+Gin project which is a result of my studying in IT Academy. </p> <p> What is it doing? Its parsing an index page in your web site, looking for a links in html. </p> <p> If it's the same like a domane it's and it'll check this link later. If it's contains '@', another domain or file extension it's ignoring. </p> <p> After checking all links It's has all links of your web-site. Next step it's creating a 'Sitemap.xml' file and send it to user's e-mail. </p>")})
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
		"text":     template.HTML(text)})
}

func PostHandler(c *gin.Context) {
	mail, err := checkIn.EmailFixer(c.PostForm("email"))
	if err != nil {
		c.Data(200, "text/html; charset=utf-8", []byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	domain, err := checkIn.DomainFixer(c.PostForm("site"))
	if err != nil {
		c.Data(200, "text/html; charset=utf-8", []byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	links, err := parser.Parser(domain)
	if err != nil {
		c.Data(200, "text/html; charset=utf-8", []byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	xmlFile, err := xmlgen.StyleCreator(links, domain)
	if err != nil {
		c.Data(200, "text/html; charset=utf-8", []byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	err = xmlgen.FileWriter(domain, xmlFile)
	if err != nil {
		c.Data(200, "text/html; charset=utf-8", []byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
	//err = sender.Email(domain, mail)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	c.Data(200, "text/html; charset=utf-8", nil)
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
