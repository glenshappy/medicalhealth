package main

import (
	"fmt"
	"strings"
)

func square(n int)  int{
	return n*n
}

func negative(n int)  int{
	return -n
}

func product(m,n int)  int{
	return m*n
}

func main()  {
	f:=square
	fmt.Println(f(3))
	f=negative
	fmt.Println(f(3))
	fmt.Printf("%T\n",f)
	str:=strings.Map(func(r rune) rune {
		return r+1
	},"H1L")
	fmt.Println(str)

}
