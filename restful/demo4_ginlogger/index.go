package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	count := 0
	runningDit, _ := os.Getwd()
	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDit), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDit), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	//r.Use(gin.Logger())
	/*r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage)
	}))*/

	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/profile"))

	r.GET("/", func(c *gin.Context) {
		c.Data(200, "text/html/charset=utf-8", []byte("Root"))
	})
	r.GET("/profile", func(c *gin.Context) {
		count = count + 1
		accesslogfile.WriteString(fmt.Sprintf("Count : %d", count))
		c.Data(200, "text/html/charset=utf-8", []byte("profile"))
	})

	r.GET("/error", func(c *gin.Context) {
		count = count + 1
		errlogfile.WriteString(fmt.Sprintf("Count : %d", count))
		c.Data(200, "text/html/charset=utf-8", []byte("error"))
	})

	r.Run(":82")
}
