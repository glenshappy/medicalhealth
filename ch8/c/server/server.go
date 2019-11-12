package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func echo(w io.Writer,s string,t time.Duration)  {
	fmt.Fprintln(w,strings.ToUpper(s))
	time.Sleep(t)
	fmt.Fprintln(w,s)
	time.Sleep(t)
	fmt.Fprintln(w,strings.ToLower(s))
}

func handleConn(conn net.Conn)  {
	//这里循环输出一个时间
	//进行输出回显
	//获取客户端输入
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan() {
		go echo(conn,scanner.Text(),time.Second)
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
