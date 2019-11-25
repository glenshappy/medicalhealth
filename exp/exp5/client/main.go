package main

import (
	"medicalhealth/exp/exp5/service"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService()  (HelloServiceClient,error){
	conn,err:=net.Dial("tcp",":1234")
	if err!=nil {
		return HelloServiceClient{},err
	}
	c:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return HelloServiceClient{c},nil
}

func (p *HelloServiceClient) Hello(request string,reply *string) error {
	return p.Client.Call(service.ServiceName+".Hello",request,reply)
}

var wg sync.WaitGroup
func main() {
	hc,err:=DialHelloService()
	if err!=nil {
		panic(err)
	}
	var reply string
	err=hc.Hello(" wanxh",&reply)
	if err!=nil {
		panic(err)
	}
	println(reply)
}
