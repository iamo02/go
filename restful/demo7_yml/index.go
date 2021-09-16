package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hay", "status": http.StatusOK})
	})

	r.GET("moreJson", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Pongchai"
		msg.Message = "Hey"
		msg.Number = 123

		c.JSON(http.StatusOK, msg)

	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hay", "status": http.StatusOK})
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hay", "status": http.StatusOK})
	})

	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "string")
	})

	r.Run(":82")
}
