package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

var(
	cl ConnLimiter
)

type middleWareHandler struct {
	r *httprouter.Router
}

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/uploadVideo",uploadVideo)
	router.GET("/video/:vid",getVideo)
	router.GET("/testpage",testPageHandler)
	return router
}

func NewMiddleWareHandler()  middleWareHandler {
	mh:=middleWareHandler{}
	mh.r = RegisterHandlers()
	return mh
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//实现限流
	rs:=cl.GetConn()
	if(rs!=true){
		sendErroMsg(w,http.StatusTooManyRequests,"too many requests")
		return
	}
	m.r.ServeHTTP(w,r)
	defer cl.ReleaseConn()
}

func main() {
	cl=NewConnLimiter(2)
	mh:=NewMiddleWareHandler()
	http.ListenAndServe(":8081", mh)
}