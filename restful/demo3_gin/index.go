package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// type LoginForm struct {
// 	Username string `form:"username" binding:"required"`
// 	Password string `form:"password" binding:"required"`
// }

func handleBookReq(c *gin.Context) {
	from, to := c.Param("from"), c.Param("to")
	c.JSON(http.StatusOK, gin.H{"result": "ok", "from": from, "to": to})
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("hello"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("profile"))
	})

	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "usernsme": username, "password": password})

	})

	r.GET("/book/:from/:to", handleBookReq)

	r.POST("/login", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.Username == "admin" && form.Password == "1234" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"satus": "unathorized", "username": form.Username, "password": form.Password})
			}
		} else {
			c.JSON(401, gin.H{"status": "unble to bind"})
		}
	})

	r.Run(":82")
}
