package main

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"medicalhealth/exp/exp5/proto"
)

func init()  {
	generator.RegisterPlugin(new(proto.NetrpcPlugin))
}

func main() {
	
}
