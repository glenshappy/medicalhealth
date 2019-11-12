package main1

import (
	"flag"
	"fmt"
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

func main() {
	var n,n1 string
	flag.StringVar(&n,"n","ss","input n")
	flag.StringVar(&n1,"s", " ", "separator")
	n2:=flag.String("n2","n2_default","input n2:")
	flag.Parse()
	fmt.Println(n,n1,*n2)
	//fmt.Print(strings.Join(flag.Args(), *sep))
	//if !*n {
	//	fmt.Println()
	//}
}