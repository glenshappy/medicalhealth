package main

import (
	"flag"
	"fmt"
	//"strings"
)

func incr(p *int)  int{
	*p++
	return *p
}

func newInt()  *int {
	return new(int) //返回了一个匿名变量的地址
}

func newInt1()  *int{
	var dummy int
	return &dummy
}

func delta(new int)  int{
	return new+1
}
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
func main() {
	flag.Parse()
	fmt.Println("sjsjsj:",*n,*sep)
	//fmt.Print(strings.Join(flag.Args(), *sep))
	//if !*n {
	//	fmt.Println()
	//}
}