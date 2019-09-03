package main

import (
	"fmt"
	"time"
)

func fib(n int)  int{
	if n<2 {
		return n
	}
	return fib(n-1)+fib(n-2)
}

func spinner(time1 time.Duration)  {
	for _,v:=range `\|/` {
		fmt.Printf("\r%c",v)
		time.Sleep(time1)
	}
}

func main()  {
	go spinner(1000*time.Millisecond)
	const  n  = 45
	fibN:= fib(n)
	fmt.Println("fibn:",fibN)
}
