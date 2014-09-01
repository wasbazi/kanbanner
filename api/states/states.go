package states

import (
	"github.com/gin-gonic/gin"
	"github.com/wasbazi/kanbanner/db/states"
)

func ViewAllHandler(c *gin.Context) {
	states, _ := states.LoadStates()

	c.JSON(200, states)
}

func SetupHandlers(r *gin.Engine) {
	r.GET("/states", ViewAllHandler)
}
