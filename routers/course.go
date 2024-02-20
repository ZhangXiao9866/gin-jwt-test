package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/course"
)

func InitCourse(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/course/:id", course.Get)
		v1.POST("/course", course.Add)
		v1.PUT("/course", course.Updata)
		v1.DELETE("/course", course.Delete)
	}
}
