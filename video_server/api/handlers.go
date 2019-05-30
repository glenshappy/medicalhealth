package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"medicalhealth/video_server/api/dbops"
	"medicalhealth/video_server/api/defs"
	"medicalhealth/video_server/api/sessions"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//io.WriteString(w,"create user")
	loginName:=r.PostFormValue("login_name")
	pwd:= r.PostFormValue("pwd")
	err:=dbops.AddUserCredential(loginName,pwd)
	if err!=nil{
		defs.SendJsonMsg(w,http.StatusInternalServerError,defs.ReponseMsg{Code:http.StatusInternalServerError,Msg:err.Error()})
		return
	}
	sid,err:=sessions.GenerateNewSessionId(loginName)
	if err!=nil {
		defs.SendJsonMsg(w,http.StatusInternalServerError,defs.ReponseMsg{Code:http.StatusInternalServerError,Msg:err.Error()})
	}

	defs.SendJsonMsg(w,http.StatusOK,defs.ReponseMsg{Code:200,Msg:"成功",Data: map[string]interface{}{
		"sid":sid,
	}})
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	loginName:=p.ByName("login_name")
	pwd:=p.ByName("pwd")
	pwdDb,err:=dbops.GetUserCredential(loginName)
	if err!=nil {
		defs.SendJsonMsg(w,http.StatusInternalServerError,defs.ReponseMsg{Code:http.StatusInternalServerError,Msg:err.Error()})
		return
	}
	if pwd!=pwdDb {
		defs.SendJsonMsg(w,http.StatusOK,defs.ReponseMsg{Code:201,Msg:"密码不正确"})
		return
	}
	sid,err:=sessions.GenerateNewSessionId(loginName)

	if err!=nil {
		defs.SendJsonMsg(w,http.StatusOK,defs.ReponseMsg{Code:202,Msg:"登录失败，请重试！"})
	}
	defs.SendJsonMsg(w,http.StatusOK,defs.ReponseMsg{Code:200,Msg:"登录成功",Data: map[string]interface{}{
		"sid":sid,
	}})
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