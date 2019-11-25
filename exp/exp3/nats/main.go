package main

import (
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc,err:=nats.Connect(nats.DefaultURL)
	if err!=nil {
		panic(err)
	}
	msg:="hello wanjn:"+time.Now().Format("2006/1/2 15:04:05")
	err=nc.Publish("tasks",[]byte(msg))
	if err!=nil {
		panic(err)
	}
	nc.Flush()
}
