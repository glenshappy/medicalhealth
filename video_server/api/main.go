package main

import (
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/julienschmidt/httprouter"
	"medicalhealth/video_server/api/sessions"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:login_name/:pwd", Login)
	routers.Get("/user/:sid", GetUserInfo)
	router.GET("/video/:vid", GetVideo)
	router.POST("/comment", CreateComment)
	router.GET("/comment/:id", GetComment)
	return router
}

func Prepare() {
	sessions.LoadSessionsFromDB()
}

func main() {
	Prepare()
	router := RegisterHandlers()
	http.ListenAndServe(":8080", router)
}
