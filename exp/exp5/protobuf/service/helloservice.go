package service

import (
	"net/rpc"
	"medicalhealth/exp/exp5/proto"
)

const ServiceName  = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(*proto.String,*proto.String) error
}

type HelloService struct {

}

func (p *HelloService) Hello(request *proto.String,reply *proto.String)  error{
	reply.Value="hello"+request.GetValue()
	return nil
}

func  RegisterHelloService(rcvr HelloServiceInterface) error{
	return rpc.RegisterName(ServiceName,rcvr)
}

