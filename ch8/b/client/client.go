package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func mustCopy(writer io.Writer,reader io.Reader)  {
	if _,err:=io.Copy(writer,reader);err != nil{
		log.Fatal(err)
	}
}

func handleConn(address string)  {
	conn,err:=net.Dial("tcp",address)
	if err!=nil {
		log.Fatal(err)
	}
	defer  conn.Close()
	mustCopy(os.Stdout,conn)
}

func main()  {
	addresses:=os.Args[1:]
	fmt.Println(addresses)
	//循环创建通信套接字

	for _,address:=range addresses {
		//一次显示所有的结果表示需要并发通信，此时采用携程
		//解析地址
		idx:=strings.Index(address,"=")
		addressRel:=address[idx+1:]
		go handleConn(addressRel)
	}

	time.Sleep(20*time.Second)
}
