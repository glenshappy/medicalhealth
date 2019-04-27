package form

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Student struct {
	Colors []string `form:"colors[]"`
}
func main() {
	route:=gin.Default()
	route.Any("/bindform", func(c *gin.Context) {
		var stu Student
		c.ShouldBind(&stu)
		c.JSON(http.StatusOK,gin.H{"color":stu.Colors})
	})
	route.Run(":8080")
}
