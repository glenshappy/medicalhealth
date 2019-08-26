package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(n int)  {
	defer fmt.Printf("f(%d)\n",n+0/n)
	fmt.Println(n)
	f(n-1)
}

func printStack()  {
	var buf[4096]byte
	n:=buf[:]
	runtime.Stack(n,true)
	os.Stdout.Write(n)
}

func main()  {
	//延迟函数调用在释放堆栈之前
	defer printStack() //defer入栈
	f(3)

}
