package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "add",
	})
}
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "get",
	})
}
func Updata(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "up",
	})
}
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete",
	})
}
