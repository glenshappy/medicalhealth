package main

import (
	"log"
	"medicalhealth/exp/exp5/service"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)
//基于tcp实现的golang的rpc编解码器
func main() {
	service.RegisterHelloService(new(service.HelloService))
	listener,err:=net.Listen("tcp",":1234")
	if err!=nil {
		panic(err)
	}
	for {
		conn,err:=listener.Accept()
		if err!=nil {
			log.Fatal("Accept error:",err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
