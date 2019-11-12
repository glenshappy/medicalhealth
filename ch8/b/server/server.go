package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
)


func handleConn(conn net.Conn)  {
	//这里循环输出一个时间
	for {
		_,err:=io.WriteString(conn,time.Now().Format("15:04:05\n")+"\n")
		if err!=nil {
			break
		}
		time.Sleep(time.Second)
	}
	defer conn.Close()
}


var port = flag.String("port","8090","请输入端口")
func main()  {
	flag.Parse()
	listener,err:=net.Listen("tcp","localhost:"+*port)
	if err!=nil {
		log.Fatal(err)
	}
	for {
		conn,err:=listener.Accept() //如果没有新的套接字，程序运行到这里将会阻塞
		if err!=nil {
			break
		}
		go handleConn(conn)
	}
}
