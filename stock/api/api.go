package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

func Sutup(router *gin.Engine) {
	db.SetupDB()
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTtrsnsactionAPI(router)
}
