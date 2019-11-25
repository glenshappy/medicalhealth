package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"io/ioutil"
	proto2 "medicalhealth/exp/exp5/proto"
	"os"
)

func init()  {
	generator.RegisterPlugin(new(proto2.NetrpcPlugin))
}

func main() {
	g:=generator.New()
	data,err:=ioutil.ReadAll(os.Stdin)
	if err!=nil {
		g.Error(err,"reading input")
	}
	if err:=proto.Unmarshal(data,g.Request);err!=nil {
		g.Error(err,"parsing input error")
	}

	if len(g.Request.FileToGenerate)==0 {
		g.Fail("no input file")
	}
	g.CommandLineParameters(g.Request.GetParameter())
	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.GenerateAllFiles()
	data,err=proto.Marshal(g.Response)
	if err!=nil {
		g.Error(err,"failed to marshal output proto")
	}
	_,err=os.Stdout.Write(data)
	if err!=nil {
		g.Error(err,"failed to write output proto")
	}
}
