package stories

import (
	"github.com/gin-gonic/gin"
	"github.com/wasbazi/kanbanner/db/stories"
)

func CreateHandler(c *gin.Context) {
	type Story struct {
		Title string `json:"title" bind:"required"`
		Body  string `json:"body" bind:"required"`
	}

	var story Story

	if c.EnsureBody(&story) {
		id, err := stories.CreateStory(story.Title, story.Body)

		if err != nil {
			c.JSON(404, err)
		}

		m := make(map[string]string)
		m["id"] = id
		c.JSON(200, m)
	}
}

func DeleteHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	err := stories.DeleteStory(id)

	if err != nil {
		c.JSON(404, err)
	}

	m := make(map[string]bool)
	m["deleted"] = true
	c.JSON(200, m)
}

func EditHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var story stories.Story

	if c.EnsureBody(&story) {
		stories.EditStory(id, story)
		story, err := stories.LoadStory(id)

		if err != nil {
			c.JSON(404, err)
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
	r.GET("/stories/:id", ViewHandler)
	r.POST("/stories/:id", EditHandler)
	r.DELETE("/stories/:id", DeleteHandler)
	r.POST("/stories", CreateHandler)
}
