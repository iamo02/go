package api

import (
	"github.com/gin-gonic/gin"
)

func Sutup(r *gin.Engine) {
	authenAPI := r.Group("/authen")
	{
		authenAPI.GET("/login", Login)
		authenAPI.GET("/register", Register)
	}
}
