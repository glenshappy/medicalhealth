package main

import (
	"medicalhealth/exp/exp4/rpc"
)

type Student struct {
	rpc.Xtype
}

func main() {
	s:=Student{Xtype:rpc.Xtype{"www"}}
	println(s.Xtype.Hello())
}
