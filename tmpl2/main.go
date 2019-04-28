package tmpl2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func formatAsDate(t time.Time)  string{
	year,month,day:=t.Date()
	return fmt.Sprintf("%d/%02d/%02d",year,month,day)
}

func main(){
	route:=gin.Default()
	route.Delims("{{","}}")
	route.SetFuncMap(map[string]interface{}{
		"formatAsDate":formatAsDate,
	})
	route.LoadHTMLGlob("templates/*/*")
	route.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
			"title":"posts",
			"now":time.Date(2019,4,28,16,12,39,0,time.UTC),
		})
	})
	route.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"users",
			"now":time.Date(2019,4,27,16,11,11,0,time.UTC),
		})
	})
	route.Run(":8081")
}
