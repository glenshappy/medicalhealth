package service

import (
	"net/rpc"
)

const ServiceName  = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(string,*string) error
}

type HelloService struct {
	
}

func (p *HelloService) Hello(request string,reply *string)  error{
	*reply="hello"+request
	return nil
}

func  RegisterHelloService(rcvr HelloServiceInterface) error{
	return rpc.RegisterName(ServiceName,rcvr)
}

