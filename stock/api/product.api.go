package api

import (
	"main/interceptor"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.POST("/product", crrateProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "getProduct", "username": c.GetString("jwt_username")})
}

func crrateProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "crrateProduct"})
}
