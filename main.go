package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	route:=gin.Default()
	route.GET("/some_data_form_reader", func(c *gin.Context) {
		response,err:=http.Get("https://www.baidu.com/img/bd_logo1.png")
		if err!=nil || response.StatusCode!=http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
		}
		reader:=response.Body
		contentLength:=response.ContentLength
		contentType:=response.Header.Get("Content-Type")
		//fmt.Println(contentType)
		extraHeaders := map[string]string{
			//"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	route.Run(":8081")
}
