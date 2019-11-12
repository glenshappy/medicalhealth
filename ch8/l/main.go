package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var rw sync.RWMutex
func deposit(amount int)  {
	rw.Lock()
	balaceMoney+=amount
	rw.Unlock()
}

func balance()  int{
	rw.RLock()
	defer rw.RUnlock()
	return balaceMoney
}

var balaceMoney int
var wg sync.WaitGroup
func init()  {
	balaceMoney=0
}

func main()  {
	wg.Add(3)
	go func() {
		deposit(10)
		wg.Done()
	}()
	go func() {
		deposit(20)
		wg.Done()
	}()
	go func() {
		deposit(90)
		wg.Done()
	}()

	bal:=balance()
	fmt.Println("余额1：",bal)
	wg.Wait()
	for i:=0;i<100;i++ {
		bal=balance()
		fmt.Println("余额2：",bal)
	}

}
