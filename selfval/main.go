package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var controller chan string
var ticker *time.Ticker

func NewController()  (chan string){
	controller = make(chan  string,3)
	return controller
}

func WorkerStart()  {
	for {
		select {
			case <-ticker.C:
				fmt.Println("每隔3秒tick一次")
				//go func( ) {
				//	//这是一个dispatch和executer的携程
				//ForLoop:
				//	for  {
				//		select {
				//		case c:=<-controller:
				//			fmt.Println(c)
				//		default:
				//			break ForLoop
				//		}
				//	}
				//}()
		default:

		}
	}
}

func main()  {
	NewController()
	ticker = time.NewTicker(3*time.Second)
	go WorkerStart()
	c:=make(chan os.Signal,1)
	signal.Notify(c,syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s:=<-c
	fmt.Println("信号：",s)
	switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:

		default:
			return
	}
}