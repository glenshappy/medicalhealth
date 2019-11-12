package main

import (
	"io"
	"log"
	"net"
	"os"
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
	go mustCopy(os.Stdout,conn)
	mustCopy(conn,os.Stdin)
}

func main()  {

	handleConn("localhost:8090")
	time.Sleep(20*time.Second)
}
