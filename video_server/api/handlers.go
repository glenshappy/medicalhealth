package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"medicalhealth/video_server/api/dbops"
	"medicalhealth/video_server/api/defs"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//io.WriteString(w,"create user")

}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	user_name:=p.ByName("user_name")
	io.WriteString(w,user_name)
}

func GetVideo(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	vid:=p.ByName("vid")
	videoOne,err:=dbops.GetVideoInfo(vid)
	if err!=nil {
		defs.SendErrorMsg(w,http.StatusNotFound,"没找到数据")
	}
	dataStr,_:= json.Marshal(videoOne)
	fmt.Println(videoOne)
	io.WriteString(w,string(dataStr))
}

func CreateComment(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

}

func GetComment(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

}