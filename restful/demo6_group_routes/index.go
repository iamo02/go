package main

import (
	"main/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api.Sutup(r)

	stockAPI := r.Group("/stock")
	{
		stockAPI.GET("/list", listProduct)
		stockAPI.GET("/proguct", createProduct)
	}

	r.Run(":82")
}

func listProduct(c *gin.Context) {
	c.String(http.StatusOK, "List Product")
}

func createProduct(c *gin.Context) {
	c.String(http.StatusOK, "Create Product")
}
