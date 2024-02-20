package routers

import (
	"main.go/user"

	"github.com/gin-gonic/gin"
)

func InitUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/user/:id", user.Get)
		v1.POST("/user", user.Add)
		v1.PUT("/user", user.Updata)
		v1.DELETE("/user", user.Delete)
	}
	v2 := group.Group("/v2")
	{
		v2.GET("/user/:id", user.Get_v2)
		v2.POST("/user", user.Add_v2)

	}
}
