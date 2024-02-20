package main

import (
	"github.com/gin-gonic/gin"
	"main.go/routers"
)

func main() {
	r := gin.Default()
	//v1 := r.Group("/v1")
	//v2 := r.Group("/v2")
	routers.InitRouters(r)

}
