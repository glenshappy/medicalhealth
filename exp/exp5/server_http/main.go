package main

import (
	"io"
	"medicalhealth/exp/exp5/service"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)
//基于http实现的golang的rpc编解码器

func main() {
	service.RegisterHelloService(new(service.HelloService))
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser:r.Body,
			Writer:w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234",nil)
}
