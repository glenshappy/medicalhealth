package main

import (
	"net"
	"strings"
)

//返回接受的信息，小写转成大写字母
func replyConn(c net.Conn) error {
	data,err:= pools.Read(c)
	if err!=nil{
		return err
	}
	err= pools.Write(c,[]byte(strings.ToUpper(string(data))))
	return err
}

func main()  {
	ln, err := net.Listen("tcp", "127.0.0.1:8972")
	if err != nil {
		panic(err)
	}
	epoller, err = MkEpoll()
	if err != nil {
		panic(err)
	}

	go start()

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		if err := epoller.Add(conn); err != nil {
			log.Printf("failed to add connection %v", err)
			conn.Close()
		}
	}
}