package jsonp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.Any("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"color": "red",
		}
		c.JSONP(http.StatusOK, data)
	})
	route.Static("/assets2", "./assets")
	route.StaticFS("/more_static", http.Dir("my_file_system"))
	route.StaticFile("/favicon.ico", "./resources/favicon.ico")
	route.Run(":8080")
}
