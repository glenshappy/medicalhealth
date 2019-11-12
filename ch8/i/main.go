package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"runtime"
)

//每一个和客户端通信的channel
type client chan string

var (
	messages = make(chan string) //所有客户端发送的消息
	clients = make(map[client]bool) //声明一个map，表示全局当前存在的client有哪些
	entering = make(chan client)
	leaving = make(chan client)
)

func main()  {
	serverConn,err:=net.Listen("tcp","localhost:8888")
	if err!=nil {
		log.Fatal(err)
	}
	go broadCast()
	for{
		conn,err:=serverConn.Accept()
		if err!=nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

//处理消息广播
func broadCast()  {
	for{
		select {
			case cli:=<-entering:
				clients[cli] = true
			case cli:=<-leaving:
				delete(clients,cli)
			case msg:=<-messages:
				//此时向所有客户端的channel发送消息
				for cli,_:=range clients {
					cli<-msg
				}
		}
	}
}

//该携程用于全局通信使用
func clientWriter(conn net.Conn,ch <-chan string)  {
	for  msg:=range ch { //该消息为全局消息，来源为所有的客户端
		fmt.Fprintln(conn,msg)
	}
}

func handleConn(conn net.Conn)  {
	osType:=runtime.GOOS
	//处理每一连接
	ch:=make(chan string) //用于接收全局的消息，另外和客户端通信的channel
	go clientWriter(conn,ch)
	entering<-ch //此时将该条通信channel加入了全局clients
	//写全局消息
	who:=conn.RemoteAddr().String()+" is comming!"
	messages<-who
	//从客户端读取消息
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan() {//得到这个scanner，对于windows，结束符不是ctrl+D(d)，也不是ctrl+z，需要手动判断
		txt:=scanner.Text()
		if txt=="ret" && osType=="windows" {
			break
		}
		messages<-conn.RemoteAddr().String()+":"+txt
	}
	//向leavingMsg
	leavingMsg:=conn.RemoteAddr().String()+" is leaving!"
	messages<-leavingMsg
	leaving<-ch
	conn.Close()
}
