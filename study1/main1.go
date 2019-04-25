package study1

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)
var goos=runtime.GOOS

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/someGet/:name", func(c *gin.Context) {
		name:=c.Param("name")
		c.String(http.StatusOK,"hello %s",name)
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name:=c.Param("name")
		action:=c.Param("action")
		message:=name+" is "+action
		c.String(http.StatusOK,message)
	})
	
	r.GET("/welcome", func(c *gin.Context) {
		firstname:=c.DefaultQuery("firstname","Guest")
		lastname:=c.Query("lastname")
		c.String(http.StatusOK,"hello %s %s",firstname,lastname)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message:=c.PostForm("message")
		nick:=c.DefaultPostForm("nick","人性原罪")
		c.JSON(http.StatusOK,gin.H{
			"status":"posted",
			"message":message,
			"nick":nick,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		id:=c.Query("id")
		page:=c.DefaultQuery("page","1")
		name:=c.PostForm("name")
		message:=c.DefaultPostForm("message","this is great")
		fmt.Printf("id:%s;page:%s;name:%s;message:%s",id,page,name,message)
	})
	


	r.DELETE("/someDel", func(c *gin.Context) {
		c.String(http.StatusOK,"del")
	})
	r.PATCH("/somePatch", func(c *gin.Context) {
		c.String(http.StatusOK,"patch方法")
	})
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		//c.String(http.StatusOK, "pong")
		c.JSON(http.StatusOK,gin.H{"message":"hello"})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	r.POST("/upload", func(c *gin.Context) {
		file,err:=c.FormFile("file")
		if err != nil {
			//c.String(http.StatusBadRequest, err)
			fmt.Println(err)
			return
		}
		//log.Println(file.Filename)
		err=c.SaveUploadedFile(file,file.Filename)
		if err!=nil{
			c.String(http.StatusBadRequest,fmt.Sprintf("upload file error:%s",err.Error()))
			return ;
		}
		c.String(http.StatusOK,fmt.Sprintf("aabbcc:%s uploaded",file.Filename))
	})

	r.POST("/upload_multi", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]
		fmt.Println("sjsjsjsjs------------slsls")
		fmt.Println(form.File)
		fmt.Println("uuuuuuuuu----------sjsjs2222")
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
	})

	v1:=r.Group("/v1")
	v1.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK,"v1:login")
	})
	v1.POST("/submit", func(c *gin.Context) {
		c.String(http.StatusOK,"v1:submit")
	})
	v1.POST("/read", func(c *gin.Context) {
		c.String(http.StatusOK,"v1:read")
	})
	v2:=r.Group("/v2")
	v2.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK,"v2:login")
	})
	v2.POST("/submit", func(c *gin.Context) {
		c.String(http.StatusOK,"v2:submit")
	})
	v2.POST("/read", func(c *gin.Context) {
		c.String(http.StatusOK,"v2:read")
	})
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
