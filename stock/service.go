package main

import (
	"fmt"
	"main/api"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	r := gin.Default()

	// Set up CORS middleware options
	config := cors.Config{
		Origins:        "*",
		RequestHeaders: "Origin, Authorization, Content-Type",

		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	r.Use(cors.Middleware(config))

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
