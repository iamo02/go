package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	runningDir, _ := os.Getwd()
	count := 0

	router.MaxMultipartMemory = 8 << 20
	fmt.Println(count)
	router.POST("/upload", func(c *gin.Context) {

		count = count + 1
		fmt.Println(count)

		file, _ := c.FormFile("file")
		token := c.PostForm("token")
		username := c.PostForm("username")

		extension := filepath.Ext(file.Filename)

		c.SaveUploadedFile(file, fmt.Sprintf("%s/uploaded/%d%s", runningDir, count, extension))
		c.String(http.StatusOK, fmt.Sprintf("'%s' - [username: %s,token:%s] uploaded!", file.Filename, username, token))

	})

	router.Run(":82")
}
