package main

import (
	"github.com/labstack/gommon/log"
)

type ConnLimiter struct {
	ConcurrentConn int
	Bucket chan int
}

func NewConnLimiter(concurrentConn int)  ConnLimiter{
	connLimiter:=ConnLimiter{
		ConcurrentConn:concurrentConn,
		Bucket:make(chan  int,concurrentConn),
	}
	return connLimiter
}

func (cl *ConnLimiter) GetConn()  bool{
	if len(cl.Bucket) >=cl.ConcurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.Bucket<-1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c:=<-cl.Bucket
	log.Printf("New connction coming: %d", c)
}
