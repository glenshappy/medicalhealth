package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
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
	
}

func CreateComment(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

}

func GetComment(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

}