package main

import (
	"fmt"
	"sync"
	"time"
)

var balanceMoney int

func init()  {
	balanceMoney=0
}

func balance()  int{
	return <-balanceCh
}

//存钱
func deposit(amount int)  {
	amountCh<-amount
}

func teller() {
	for {
		select {
			case amount:=<-amountCh:
				balanceMoney+=amount
			case balanceCh<-balanceMoney:
		}
	}
}

var balanceCh chan int
var amountCh chan int
var wg sync.WaitGroup
func main()  {
	balanceCh=make(chan int)
	amountCh=make(chan  int,3)
	go teller()
	go func() {
		wg.Add(1)
		deposit(10)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		deposit(20)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		deposit(30)
		wg.Done()
	}()
	wg.Wait()

	time.Sleep(3*time.Second)
	for x:=range balanceCh{
		fmt.Println("余额1：",x)
	}

	//x:=<-balanceCh
	//fmt.Println("余额1：",x)
	//x=<-balanceCh
	//fmt.Println("余额2：",x)
	//x=<-balanceCh
	//fmt.Println("余额3：",x)

}
