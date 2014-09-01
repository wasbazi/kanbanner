package stories

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wasbazi/kanbanner/db/stories"
)

func EditHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var story stories.Story

	if c.EnsureBody(&story) {
		stories.EditStory(id, story)
		story, err := stories.LoadStory(id)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, story)
	}

}

func ViewHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	story, _ := stories.LoadStory(id)

	c.JSON(200, story)
}

func ViewAllHandler(c *gin.Context) {
	stories, _ := stories.LoadStories()

	c.JSON(200, stories)
}

func SetupHandlers(r *gin.Engine) {
	r.GET("/stories", ViewAllHandler)
	r.GET("/story/:id", ViewHandler)
	r.POST("/story/:id", EditHandler)
}
