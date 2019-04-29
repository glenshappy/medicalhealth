package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-log"
	"net/http"
	"strconv"
)

type News struct {
	title string
	content string
	description string
	url string
	category_id int
	created_at string
	updated_at string
}
var db gorm.DB

func main(){
	db, err := gorm.Open("mysql", "root:root@/medicalhealth?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	route:=gin.Default()
	route.LoadHTMLGlob("templates/*")
	route.GET("/index", func(c *gin.Context) {
		Db:=db
		page,_:=strconv.Atoi(c.DefaultQuery("page","1"))
		pageSize,_:= strconv.Atoi(c.DefaultQuery("pagesize","10"))
		Db=Db.Limit(pageSize).Offset((page-1)*pageSize)
		news:=make([]News,0)
		if err:=Db.Find(&news).Error;err!=nil {
			log.Fatal(err)
			return
		}
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"news_list":news,
		})
	})
	route.GET("/redirect1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})
	route.Run(":8081")
}
