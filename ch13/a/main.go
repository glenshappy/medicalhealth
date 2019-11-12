package main

import (
	"fmt"
	"reflect"
)

func main() {
	//a:=2
	//x:=&a
	//fmt.Println(*x)

	x:=2
	a:=reflect.ValueOf(2)
	b:=reflect.ValueOf(x)
	c:=reflect.ValueOf(&x)
	d:=c.Elem()
	px:=d.Addr().Interface().(*int)
	fmt.Println(a,b,c,d,*px)
}
