package defs

import (
	"io"
	"net/http"
)

func SendErrorMsg(w http.ResponseWriter,sc int,msg string)  {
	w.WriteHeader(sc)
	io.WriteString(w,msg)
}