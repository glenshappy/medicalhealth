package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	"unsafe"
)

type Student struct {
	a bool
	b int16
	c []int
}

func main() {
	dd:= '\U00004e16'
	fmt.Println(string(dd))

	b:=true
	fmt.Println(unsafe.Sizeof(b))
	d:=Student{true,22,[]int{1,3,4,5,6,7,8,9,10},}
	fmt.Println(unsafe.Sizeof(d))

	fmt.Println("============")
	a:=32
	fmt.Printf("%#v,%T",a,a)
	wula:=string(65)
	hehe:=strconv.FormatInt(int64(65),10)
	fmt.Printf("%#v\n",hehe)
	fmt.Printf("%#v\n",wula)

	str:="们"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	for i, r := range "Hello，世界界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
