package defs

import (
	"encoding/json"
	"io"
	"net/http"
)

func SendErrorMsg(w http.ResponseWriter,sc int,msg string)  {
	w.WriteHeader(sc)
	io.WriteString(w,msg)
}

func SendJsonMsg(w http.ResponseWriter,sc int,data interface{})  {
	w.WriteHeader(sc)
	w.Header().Set("Content-Type","application/json")
	jsonSlice,err:=json.Marshal(data)
	if err!=nil {
		panic(err)
	}
	io.WriteString(w,string(jsonSlice))
}