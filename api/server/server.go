package server

import (
	"log"

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

func AcceptConnections() {
	r := gin.Default()
	r.GET("/stories", ViewAllHandler)
	r.GET("/story/:id", ViewHandler)
	r.POST("/story/:id", EditHandler)

	r.Run(":8080")

	// http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/edit/", editHandler)
	// // http.HandleFunc("/save/", saveHandler)
	// http.ListenAndServe(":8080", nil)
}
