package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc,err:=nats.Connect(nats.DefaultURL)
	if err!=nil {
		panic(err)
	}
	sub,err:=nc.QueueSubscribeSync("tasks","works")
	if err!=nil {
		panic(err)
	}
	var msg *nats.Msg

	for {
		msg,err=sub.NextMsg(time.Hour*1000) //1000个小时没有收到消息
		if err!=nil {
			break
		}
		fmt.Println(string(msg.Data))
	}

}
