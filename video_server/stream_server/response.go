package main

import (
	"io"
	"net/http"
)

func sendErroMsg(w http.ResponseWriter, sc int, msg string) {
	w.WriteHeader(sc)
	io.WriteString(w, msg)
}
