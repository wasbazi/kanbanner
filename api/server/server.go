package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wasbazi/kanbanner/db"
)

func EditHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var story db.Story

	if c.EnsureBody(&story) {
		db.EditStory(id, story)
		story, err := db.LoadStory(id)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, story)
	}

}

func ViewHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	story, _ := db.LoadStory(id)

	c.JSON(200, story)
}

func ViewAllHandler(c *gin.Context) {
	stories, _ := db.LoadStories()

	c.JSON(200, stories)
}

func IndexHandler(c *gin.Context) {
	fileServer := http.FileServer(http.Dir("./public"))
	fileServer.ServeHTTP(c.Writer, c.Request)
}

func AcceptConnections() {
	db.GetDB()

	r := gin.Default()
	r.GET("/stories", ViewAllHandler)
	r.GET("/story/:id", ViewHandler)
	r.POST("/story/:id", EditHandler)

	r.GET("/", IndexHandler)
	r.Static("/javascript", "./public/javascript/")
	r.Static("/css", "./public/css/")

	port := ":8090"
	fmt.Printf("Running on %s", port)
	r.Run(port)
}
