package main

import (
	"fmt"
	"main/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/images", "./uploaded/images")
	api.Sutup(r)
	//	r.Run(":81")

	//	In case of running on Heroku
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("Running on Heroku using random PORT")
		r.Run()
	} else {
		fmt.Println("Environment Port : " + port)
		r.Run(":" + port)
	}
}
