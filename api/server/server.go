package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wasbazi/kanbanner/api/states"
	"github.com/wasbazi/kanbanner/api/stories"
	"github.com/wasbazi/kanbanner/db/db"
)

func IndexHandler(c *gin.Context) {
	fileServer := http.FileServer(http.Dir("./public"))
	fileServer.ServeHTTP(c.Writer, c.Request)
}

func AcceptConnections() {
	db.GetDB()

	r := gin.Default()

	stories.SetupHandlers(r)
	states.SetupHandlers(r)

	r.GET("/", IndexHandler)
	r.Static("/javascript", "./public/javascript/")
	r.Static("/css", "./public/css/")

	port := ":8090"
	fmt.Printf("Running on %s", port)
	r.Run(port)
}
