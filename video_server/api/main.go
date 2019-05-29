package main

import (
	"medicalhealth/video_server/api/sessions"
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

func Prepare()  {
	sessions.LoadSessionsFromDB()
}

func main() {
	Prepare()
	router:=RegisterHandlers()
	http.ListenAndServe(":8080", router)
}

