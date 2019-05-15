package main

import (
	"io"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	io.WriteString(w,"create user")
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	user_name:=p.ByName("user_name")
	io.WriteString(w,user_name)
}