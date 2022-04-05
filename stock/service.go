package main

import (
	"main/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/images", "./uploaded/images")
	api.Sutup(r)
	r.Run(":81")
}
