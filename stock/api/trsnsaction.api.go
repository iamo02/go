package api

import "github.com/gin-gonic/gin"

func SetupTtrsnsactionAPI(router *gin.Engine) {
	authAPI := router.Group("/api/v2")
	{
		authAPI.GET("/trsnsaction", getTrsnsaction)
		authAPI.POST("/trsnsaction", crrateTrsnsaction)
	}
}

func getTrsnsaction(c *gin.Context) {
	c.JSON(200, gin.H{"result": "getTrsnsaction"})
}

func crrateTrsnsaction(c *gin.Context) {
	c.JSON(200, gin.H{"result": "crrateTrsnsaction"})
}
