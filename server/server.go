package server

import (
	//"fmt"
	//"github.com/VitaliiMichailovich/GGSMG/parser"
	//"github.com/VitaliiMichailovich/GGSMG/uri"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
)

var Router *gin.Engine

//func LinkHandler(w http.ResponseWriter, r *http.Request) {
//	url, err := uri.URI(r.RequestURI[6:])
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	parser.Links[url] = parser.Link{Count: 1, Checked: false}
//	fmt.Println("Start:")
//	_, err = parser.Parser(url)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	//showMeLinks()
//	w.Write([]byte("OK"))
//
//}

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
		"text":     "<p> I just trying to improve my GoLang skills. </p>"})
}

func ContactHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.HTML(http.StatusOK, "pages.html", gin.H{
		"title":    "Contact",
		"subtitle": "Contact",
		"text":     "<p> admin@micro.pp.ua </p>"})
}

func PostHandler(c *gin.Context) {
	x, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%s", string(x))
	c.HTML(http.StatusOK, "pages.html", gin.H{
		"title":    "g",
		"subtitle": "g",
		"text":     x})
}

func Server() {
	// Handle the index route
	Router.GET("/", IndexHandler)
	Router.GET("/p/", ProjectHandler)
	Router.GET("/i/", AboutMeHandler)
	Router.GET("/c/", ContactHandler)
	Router.POST("/p/", PostHandler)

	//http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	//http.HandleFunc("/", IndexHandler)
	//http.HandleFunc("/link/", HandleLink)
	////port := os.Getenv("PORT")
	////http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":7777", nil)
}
