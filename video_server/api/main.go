package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/video/:vid",GetVideo)
	router.POST("/comment",CreateComment)
	router.GET("/comment/:id",GetComment)
	return router
}

func main() {
	router:=RegisterHandlers()
	http.ListenAndServe(":8080", router)
}

