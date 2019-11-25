package main

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init()  {
	generator.RegisterPlugin()
}

type xtype int

type int1 interface {
	Hello() int
}

type Student struct {
	xtype
}

func (x xtype)Hello()  int{
	return int(x)+1
}

func testInternal(x xtype)  {
	fmt.Println(x)
}

func main() {
	//m:=make([]int,10,20)
	//for i:=0;i<len(m) ; i++ {
	//	m[i]=i
	//}
	//fmt.Println(m,len(m),cap(m))
	//m=append(m,11,12,13,14,15,16,17,18,19,20,21,22)
	//fmt.Println(m,len(m),cap(m))


	//s:="中a"
	//fmt.Println("====",s[:3],"\n",s[3:],"===")
	//
	//
	//fmt.Println(string(s[3]))
	//h:=md5.New()
	//io.WriteString(h,"sjsjsjs")
	//fmt.Printf("%x",h.Sum(nil))

var int11 int1
var x xtype

x = 444

int11 = x

fmt.Println(x)

	switch int11.(type) {
	case nil:
		fmt.Println("null类型")
	case xtype:
		fmt.Println("xtype类型")
	default:
		fmt.Println("未知类型")
	}

var s int

s=222
if s>100 {
	fmt.Println("大于100")
}else if x>40 {
	fmt.Println("大于40")
}else{
	fmt.Println("what")
}

s1:=Student{
	xtype(233),
}


}
