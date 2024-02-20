package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/login"
)

func Initlogin(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/login", login.Get)

	}
}
