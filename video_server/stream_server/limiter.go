package main

import (
	"github.com/labstack/gommon/log"
)

type ConnLimiter struct {
	ConcurrentConn int
	Bucket         chan int
}

func NewConnLimiter(concurrentConn int) *ConnLimiter {
	connLimiter := &ConnLimiter{
		ConcurrentConn: concurrentConn,
		Bucket:         make(chan int, concurrentConn),
	}
	return connLimiter
}

func (cl *ConnLimiter) GetConn() bool {
	//log.Printf("lenbucket:%d-----lencon:%d\n",len(cl.Bucket),cl.ConcurrentConn)
	if len(cl.Bucket) >= cl.ConcurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	//log.Printf("开始获取：%v",time.Now().String())
	cl.Bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	//log.Printf("开始释放：%v",time.Now().String())
	c := <-cl.Bucket
	log.Printf("New connction coming: %d", c)
}
