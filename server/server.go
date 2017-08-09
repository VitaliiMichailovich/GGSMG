package server

import (
	//"fmt"
	//"github.com/VitaliiMichailovich/Sitemap/parser"
	//"github.com/VitaliiMichailovich/Sitemap/uri"
	"net/http"
	"github.com/gin-gonic/gin"
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
		"payload": "---"})
}

func ProjectHandler(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "ProjectHandler",
		"payload": "---"})
}

func Server() {
	// Handle the index route
	Router.GET("/", IndexHandler)
	Router.GET("/p/", ProjectHandler)
	Router.GET("/i/", IndexHandler)
	Router.GET("/c/", IndexHandler)
	Router.POST("/g/", IndexHandler)

	//http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	//http.HandleFunc("/", IndexHandler)
	//http.HandleFunc("/link/", HandleLink)
	////port := os.Getenv("PORT")
	////http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":7777", nil)
}
