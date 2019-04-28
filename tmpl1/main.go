package tmpl1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	route:=gin.Default()
	route.LoadHTMLGlob("templates/*")
	route.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"我的第一个模板",
		})
	})
	route.Run(":8081")
}
