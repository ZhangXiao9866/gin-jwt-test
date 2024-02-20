package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type courses struct {
	Name     string `form:"name" json:"name"`
	Teatcher string `form:"teatcher" json:"teatcher"`
	Duration int    `form:"duration" json:"duration"`
}

func Add(c *gin.Context) {
	req := &courses{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.JSON(http.StatusOK, req)
	}

}
func Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
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
