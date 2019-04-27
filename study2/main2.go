package study2

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name string `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}
func main()  {
	var route = gin.Default()
	route.Any("/test11", func(c *gin.Context) {
		var per = &Person{}
		c.ShouldBind(per)
		c.JSON(http.StatusOK,per)
	})
	route.Run(":8099")
	
}