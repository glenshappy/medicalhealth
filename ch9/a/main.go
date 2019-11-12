package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch1,ch2 chan string
var done=make(chan struct{})
func main() {
	ch1=make(chan string)
	ch2 = make(chan  string)
	wg.Add(2)
	var count=0
	var count2 = 0
	go func() {
		defer  wg.Done()
		loop1:
		for {
			select {
			case <-ch1:
				count++
				//fmt.Println(a)
			case ch2<-"pong":
			case <-done:
				//关闭ch2，清空ch1
				close(ch2)
				for range ch1 {

				}
				break loop1
			}
		}

	}()

	go func() {
		defer  wg.Done()
		loop1:
		for  {
			select {
			case <-ch2:
				count2++
				//fmt.Println(a)
			case ch1<-"ping":
			case <-done:
				//关闭ch1，清空ch2
				close(ch1)
				for range ch2 {

				}
				break loop1
			}
		}
	}()



	tick:=time.After(1*time.Minute)
	go func() {
		select {
		case <-tick:
			close(done)
		}
	}()


	wg.Wait()

	fmt.Println("执行次数：",count,count2)
}
