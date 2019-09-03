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
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Category_id int    `json:"category_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

var db gorm.DB

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(39.106.127.158:3306)/medicalhealth?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	route := gin.Default()
	route.Static("/assets", "./assets")
	route.StaticFS("/templates", http.Dir("templates"))
	route.StaticFile("/favicon.ico", "./resources/favicon.ico")
	route.GET("/index/index", func(c *gin.Context) {
		Db := db
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
		Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
		news := make([]News, 0)
		if err := Db.Find(&news).Error; err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"title":     "日常养生",
			"news_list": news,
		})
	})
	route.Run(":8081")
}
