package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/uploadVideo",uploadVideo)
	router.GET("/video/:vid",getVideo)
	router.GET("/testpage",testPageHandler)
	return router
}

func NewMiddleWareHandler(connCon int)  middleWareHandler {
	mh:=middleWareHandler{}
	mh.r = RegisterHandlers()
	mh.l = NewConnLimiter(connCon)
	return mh
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//实现限流
	if !m.l.GetConn() {
		sendErroMsg(w,http.StatusTooManyRequests,"too many requests")
		return
	}
	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}

func main() {
	//fmt.Printf("长度：%d==========\n",utf8.RuneCountInString("忍者s你"))
	mh:=NewMiddleWareHandler(2)
	http.ListenAndServe(":8081", mh)
}