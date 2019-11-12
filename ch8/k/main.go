package main

import "fmt"

func main()  {
	ch:=make(chan  int)
	go func() {
		ch<-1
		ch<-2
	}()
	a:=<-ch
	ab:=<-ch
	fmt.Println(a,ab)
}